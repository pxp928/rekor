//
// Copyright 2022 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package verify

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/cyberphone/json-canonicalization/go/src/webpki.org/jsoncanonicalizer"
	"github.com/sigstore/rekor/pkg/generated/client"
	"github.com/sigstore/rekor/pkg/generated/client/tlog"
	"github.com/sigstore/rekor/pkg/generated/models"
	"github.com/sigstore/rekor/pkg/util"
	"github.com/sigstore/sigstore/pkg/signature"
	"github.com/sigstore/sigstore/pkg/signature/options"
	"github.com/transparency-dev/merkle/proof"
	"github.com/transparency-dev/merkle/rfc6962"
)

// ProveConsistency verifies consistency between an initial, trusted STH
// and a second new STH. Callers MUST verify signature on the STHs'.
func ProveConsistency(ctx context.Context, rClient *client.Rekor,
	oldSTH *util.SignedCheckpoint, newSTH *util.SignedCheckpoint, treeID string) error {
	oldTreeSize := int64(oldSTH.Size)
	switch {
	case oldTreeSize == int64(newSTH.Size):
		if !bytes.Equal(oldSTH.Hash, newSTH.Hash) {
			return errors.New("old root hash does not match STH hash")
		}
	case oldTreeSize < int64(newSTH.Size):
		consistencyParams := tlog.NewGetLogProofParamsWithContext(ctx)
		consistencyParams.FirstSize = &oldTreeSize      // Root size at the old, or trusted state.
		consistencyParams.LastSize = int64(newSTH.Size) // Root size at the new state to verify against.
		consistencyParams.TreeID = &treeID
		consistencyProof, err := rClient.Tlog.GetLogProof(consistencyParams)
		if err != nil {
			return err
		}
		var hashes [][]byte
		for _, h := range consistencyProof.Payload.Hashes {
			b, err := hex.DecodeString(h)
			if err != nil {
				return errors.New("error decoding consistency proof hashes")
			}
			hashes = append(hashes, b)
		}
		if err := proof.VerifyConsistency(rfc6962.DefaultHasher,
			oldSTH.Size, newSTH.Size, hashes, oldSTH.Hash, newSTH.Hash); err != nil {
			return err
		}
	case oldTreeSize > int64(newSTH.Size):
		return errors.New("inclusion proof returned a tree size larger than the verified tree size")
	}
	return nil

}

// VerifyCurrentCheckpoint verifies the provided checkpoint by verifying consistency
// against a newly fetched Checkpoint.
//nolint
func VerifyCurrentCheckpoint(ctx context.Context, rClient *client.Rekor, verifier signature.Verifier,
	oldSTH *util.SignedCheckpoint) (*util.SignedCheckpoint, error) {
	// The oldSTH should already be verified, but check for robustness.
	if !oldSTH.Verify(verifier) {
		return nil, errors.New("signature on old tree head did not verify")
	}

	// Get and verify against the current STH.
	infoParams := tlog.NewGetLogInfoParamsWithContext(ctx)
	result, err := rClient.Tlog.GetLogInfo(infoParams)
	if err != nil {
		return nil, err
	}

	logInfo := result.GetPayload()
	sth := util.SignedCheckpoint{}
	if err := sth.UnmarshalText([]byte(*logInfo.SignedTreeHead)); err != nil {
		return nil, err
	}

	// Verify the signature on the SignedCheckpoint.
	if !sth.Verify(verifier) {
		return nil, errors.New("signature on tree head did not verify")
	}

	// Now verify consistency up to the STH.
	if err := ProveConsistency(ctx, rClient, oldSTH, &sth, *logInfo.TreeID); err != nil {
		return nil, err
	}
	return &sth, nil
}

// VerifyInclusion verifies an entry's inclusion proof. Clients MUST either verify
// the root hash against a new STH (via VerifyCurrentCheckpoint) or against a
// trusted, existing STH (via ProveConsistency).
//nolint
func VerifyInclusion(ctx context.Context, e *models.LogEntryAnon) error {
	if e.Verification == nil || e.Verification.InclusionProof == nil {
		return errors.New("inclusion proof not provided")
	}

	hashes := [][]byte{}
	for _, h := range e.Verification.InclusionProof.Hashes {
		hb, _ := hex.DecodeString(h)
		hashes = append(hashes, hb)
	}

	rootHash, err := hex.DecodeString(*e.Verification.InclusionProof.RootHash)
	if err != nil {
		return err
	}

	// Verify the inclusion proof.
	entryBytes, err := base64.StdEncoding.DecodeString(e.Body.(string))
	if err != nil {
		return err
	}
	leafHash := rfc6962.DefaultHasher.HashLeaf(entryBytes)

	if err := proof.VerifyInclusion(rfc6962.DefaultHasher, uint64(*e.Verification.InclusionProof.LogIndex),
		uint64(*e.Verification.InclusionProof.TreeSize), leafHash, hashes, rootHash); err != nil {
		return err
	}

	return nil
}

// VerifySignedEntryTimestamp verifies the entry's SET against the provided
// public key.
//nolint
func VerifySignedEntryTimestamp(ctx context.Context, e *models.LogEntryAnon, verifier signature.Verifier) error {
	if e.Verification == nil {
		return fmt.Errorf("missing verification")
	}
	if e.Verification.SignedEntryTimestamp == nil {
		return fmt.Errorf("signature missing")
	}

	type bundle struct {
		Body           interface{} `json:"body"`
		IntegratedTime int64       `json:"integratedTime"`
		// Note that this is the virtual index.
		LogIndex int64  `json:"logIndex"`
		LogID    string `json:"logID"`
	}
	bundlePayload := bundle{
		Body:           e.Body,
		IntegratedTime: *e.IntegratedTime,
		LogIndex:       *e.LogIndex,
		LogID:          *e.LogID,
	}
	contents, err := json.Marshal(bundlePayload)
	if err != nil {
		return fmt.Errorf("marshaling bundle: %w", err)
	}
	canonicalized, err := jsoncanonicalizer.Transform(contents)
	if err != nil {
		return fmt.Errorf("canonicalizing bundle: %w", err)
	}

	// verify the SET against the public key
	if err := verifier.VerifySignature(bytes.NewReader(e.Verification.SignedEntryTimestamp),
		bytes.NewReader(canonicalized), options.WithContext(ctx)); err != nil {
		return fmt.Errorf("unable to verify bundle: %w", err)
	}
	return nil
}

// VerifyLogEntry performs verification of a LogEntry given a Rekor verifier.
// Performs inclusion proof verification up to a verified root hash and
// SignedEntryTimestamp verification.
//nolint
func VerifyLogEntry(ctx context.Context, e *models.LogEntryAnon, verifier signature.Verifier) error {
	// Verify the inclusion proof using the body's leaf hash.
	if err := VerifyInclusion(ctx, e); err != nil {
		return err
	}

	// TODO: If/when we return an STH in the response, verify that too, against an
	// optional known STH as well.
	// See https://github.com/sigstore/rekor/issues/988

	// Verify the Signed Entry Timestamp.
	if err := VerifySignedEntryTimestamp(ctx, e, verifier); err != nil {
		return err
	}

	return nil
}
