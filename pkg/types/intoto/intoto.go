//
// Copyright 2021 The Sigstore Authors.
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

package intoto

import (
	"context"
	"errors"
	"fmt"

	"github.com/sigstore/rekor/pkg/generated/models"
	"github.com/sigstore/rekor/pkg/types"
	"golang.org/x/exp/slices"
)

const (
	KIND = "intoto"
)

type BaseIntotoType struct {
	types.RekorType
}

func init() {
	types.TypeMap.Store(KIND, New)
}

func New() types.TypeImpl {
	bit := BaseIntotoType{}
	bit.Kind = KIND
	bit.VersionMap = VersionMap
	return &bit
}

var VersionMap = types.NewSemVerEntryFactoryMap()

func (it BaseIntotoType) UnmarshalEntry(pe models.ProposedEntry) (types.EntryImpl, error) {
	if pe == nil {
		return nil, errors.New("proposed entry cannot be nil")
	}

	in, ok := pe.(*models.Intoto)
	if !ok {
		return nil, errors.New("cannot unmarshal non-Rekord types")
	}

	return it.VersionedUnmarshal(in, *in.APIVersion)
}

func (it *BaseIntotoType) CreateProposedEntry(ctx context.Context, version string, props types.ArtifactProperties) (models.ProposedEntry, error) {
	if version == "" {
		version = it.DefaultVersion()
	}
	ei, err := it.VersionedUnmarshal(nil, version)
	if err != nil {
		return nil, fmt.Errorf("fetching Intoto version implementation: %w", err)
	}
	return ei.CreateFromArtifactProperties(ctx, props)
}

func (it BaseIntotoType) DefaultVersion() string {
	return "0.0.2"
}

// SupportedVersions returns the supported versions for this type;
// it deliberately omits 0.0.1 from the list of supported versions as that
// version did not persist signatures inside the log entry
func (it BaseIntotoType) SupportedVersions() []string {
	return []string{"0.0.2"}
}

// IsSupportedVersion returns true if the version can be inserted into the log, and false if not
func (it *BaseIntotoType) IsSupportedVersion(proposedVersion string) bool {
	return slices.Contains(it.SupportedVersions(), proposedVersion)
}
