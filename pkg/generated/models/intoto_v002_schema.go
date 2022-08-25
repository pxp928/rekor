// Code generated by go-swagger; DO NOT EDIT.

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
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IntotoV002Schema intoto v0.0.2 Schema
//
// Schema for intoto object
//
// swagger:model intotoV002Schema
type IntotoV002Schema struct {

	// content
	// Required: true
	Content *IntotoV002SchemaContent `json:"content"`
}

// Validate validates this intoto v002 schema
func (m *IntotoV002Schema) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntotoV002Schema) validateContent(formats strfmt.Registry) error {

	if err := validate.Required("content", "body", m.Content); err != nil {
		return err
	}

	if m.Content != nil {
		if err := m.Content.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this intoto v002 schema based on the context it is used
func (m *IntotoV002Schema) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntotoV002Schema) contextValidateContent(ctx context.Context, formats strfmt.Registry) error {

	if m.Content != nil {
		if err := m.Content.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IntotoV002Schema) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntotoV002Schema) UnmarshalBinary(b []byte) error {
	var res IntotoV002Schema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IntotoV002SchemaContent intoto v002 schema content
//
// swagger:model IntotoV002SchemaContent
type IntotoV002SchemaContent struct {

	// envelope
	Envelope *IntotoV002SchemaContentEnvelope `json:"envelope,omitempty"`

	// hash
	Hash *IntotoV002SchemaContentHash `json:"hash,omitempty"`

	// payload hash
	PayloadHash *IntotoV002SchemaContentPayloadHash `json:"payloadHash,omitempty"`
}

// Validate validates this intoto v002 schema content
func (m *IntotoV002SchemaContent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvelope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayloadHash(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntotoV002SchemaContent) validateEnvelope(formats strfmt.Registry) error {
	if swag.IsZero(m.Envelope) { // not required
		return nil
	}

	if m.Envelope != nil {
		if err := m.Envelope.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content" + "." + "envelope")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content" + "." + "envelope")
			}
			return err
		}
	}

	return nil
}

func (m *IntotoV002SchemaContent) validateHash(formats strfmt.Registry) error {
	if swag.IsZero(m.Hash) { // not required
		return nil
	}

	if m.Hash != nil {
		if err := m.Hash.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content" + "." + "hash")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content" + "." + "hash")
			}
			return err
		}
	}

	return nil
}

func (m *IntotoV002SchemaContent) validatePayloadHash(formats strfmt.Registry) error {
	if swag.IsZero(m.PayloadHash) { // not required
		return nil
	}

	if m.PayloadHash != nil {
		if err := m.PayloadHash.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content" + "." + "payloadHash")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content" + "." + "payloadHash")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this intoto v002 schema content based on the context it is used
func (m *IntotoV002SchemaContent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEnvelope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHash(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePayloadHash(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntotoV002SchemaContent) contextValidateEnvelope(ctx context.Context, formats strfmt.Registry) error {

	if m.Envelope != nil {
		if err := m.Envelope.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content" + "." + "envelope")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content" + "." + "envelope")
			}
			return err
		}
	}

	return nil
}

func (m *IntotoV002SchemaContent) contextValidateHash(ctx context.Context, formats strfmt.Registry) error {

	if m.Hash != nil {
		if err := m.Hash.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content" + "." + "hash")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content" + "." + "hash")
			}
			return err
		}
	}

	return nil
}

func (m *IntotoV002SchemaContent) contextValidatePayloadHash(ctx context.Context, formats strfmt.Registry) error {

	if m.PayloadHash != nil {
		if err := m.PayloadHash.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content" + "." + "payloadHash")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content" + "." + "payloadHash")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IntotoV002SchemaContent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntotoV002SchemaContent) UnmarshalBinary(b []byte) error {
	var res IntotoV002SchemaContent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IntotoV002SchemaContentEnvelope dsse envelope
//
// swagger:model IntotoV002SchemaContentEnvelope
type IntotoV002SchemaContentEnvelope struct {

	// payload of the envelope
	// Format: byte
	Payload strfmt.Base64 `json:"payload,omitempty"`

	// type describing the payload
	// Required: true
	PayloadType *string `json:"payloadType"`

	// collection of all signatures of the envelope's payload
	// Required: true
	// Min Items: 1
	Signatures []*IntotoV002SchemaContentEnvelopeSignaturesItems0 `json:"signatures"`
}

// Validate validates this intoto v002 schema content envelope
func (m *IntotoV002SchemaContentEnvelope) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayloadType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignatures(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntotoV002SchemaContentEnvelope) validatePayloadType(formats strfmt.Registry) error {

	if err := validate.Required("content"+"."+"envelope"+"."+"payloadType", "body", m.PayloadType); err != nil {
		return err
	}

	return nil
}

func (m *IntotoV002SchemaContentEnvelope) validateSignatures(formats strfmt.Registry) error {

	if err := validate.Required("content"+"."+"envelope"+"."+"signatures", "body", m.Signatures); err != nil {
		return err
	}

	iSignaturesSize := int64(len(m.Signatures))

	if err := validate.MinItems("content"+"."+"envelope"+"."+"signatures", "body", iSignaturesSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.Signatures); i++ {
		if swag.IsZero(m.Signatures[i]) { // not required
			continue
		}

		if m.Signatures[i] != nil {
			if err := m.Signatures[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("content" + "." + "envelope" + "." + "signatures" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("content" + "." + "envelope" + "." + "signatures" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this intoto v002 schema content envelope based on the context it is used
func (m *IntotoV002SchemaContentEnvelope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSignatures(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntotoV002SchemaContentEnvelope) contextValidateSignatures(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Signatures); i++ {

		if m.Signatures[i] != nil {
			if err := m.Signatures[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("content" + "." + "envelope" + "." + "signatures" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("content" + "." + "envelope" + "." + "signatures" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IntotoV002SchemaContentEnvelope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntotoV002SchemaContentEnvelope) UnmarshalBinary(b []byte) error {
	var res IntotoV002SchemaContentEnvelope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IntotoV002SchemaContentEnvelopeSignaturesItems0 a signature of the envelope's payload along with the public key for the signature
//
// swagger:model IntotoV002SchemaContentEnvelopeSignaturesItems0
type IntotoV002SchemaContentEnvelopeSignaturesItems0 struct {

	// optional id of the key used to create the signature
	Keyid string `json:"keyid,omitempty"`

	// public key that corresponds to this signature
	// Read Only: true
	// Format: byte
	PublicKey strfmt.Base64 `json:"publicKey,omitempty"`

	// signature of the payload
	// Format: byte
	Sig strfmt.Base64 `json:"sig,omitempty"`
}

// Validate validates this intoto v002 schema content envelope signatures items0
func (m *IntotoV002SchemaContentEnvelopeSignaturesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this intoto v002 schema content envelope signatures items0 based on the context it is used
func (m *IntotoV002SchemaContentEnvelopeSignaturesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePublicKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntotoV002SchemaContentEnvelopeSignaturesItems0) contextValidatePublicKey(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "publicKey", "body", strfmt.Base64(m.PublicKey)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IntotoV002SchemaContentEnvelopeSignaturesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntotoV002SchemaContentEnvelopeSignaturesItems0) UnmarshalBinary(b []byte) error {
	var res IntotoV002SchemaContentEnvelopeSignaturesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IntotoV002SchemaContentHash Specifies the hash algorithm and value encompassing the entire signed envelope
//
// swagger:model IntotoV002SchemaContentHash
type IntotoV002SchemaContentHash struct {

	// The hashing function used to compute the hash value
	// Required: true
	// Enum: [sha256]
	Algorithm *string `json:"algorithm"`

	// The hash value for the archive
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this intoto v002 schema content hash
func (m *IntotoV002SchemaContentHash) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var intotoV002SchemaContentHashTypeAlgorithmPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sha256"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		intotoV002SchemaContentHashTypeAlgorithmPropEnum = append(intotoV002SchemaContentHashTypeAlgorithmPropEnum, v)
	}
}

const (

	// IntotoV002SchemaContentHashAlgorithmSha256 captures enum value "sha256"
	IntotoV002SchemaContentHashAlgorithmSha256 string = "sha256"
)

// prop value enum
func (m *IntotoV002SchemaContentHash) validateAlgorithmEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, intotoV002SchemaContentHashTypeAlgorithmPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IntotoV002SchemaContentHash) validateAlgorithm(formats strfmt.Registry) error {

	if err := validate.Required("content"+"."+"hash"+"."+"algorithm", "body", m.Algorithm); err != nil {
		return err
	}

	// value enum
	if err := m.validateAlgorithmEnum("content"+"."+"hash"+"."+"algorithm", "body", *m.Algorithm); err != nil {
		return err
	}

	return nil
}

func (m *IntotoV002SchemaContentHash) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("content"+"."+"hash"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this intoto v002 schema content hash based on the context it is used
func (m *IntotoV002SchemaContentHash) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *IntotoV002SchemaContentHash) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntotoV002SchemaContentHash) UnmarshalBinary(b []byte) error {
	var res IntotoV002SchemaContentHash
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IntotoV002SchemaContentPayloadHash Specifies the hash algorithm and value covering the payload within the DSSE envelope
//
// swagger:model IntotoV002SchemaContentPayloadHash
type IntotoV002SchemaContentPayloadHash struct {

	// The hashing function used to compute the hash value
	// Required: true
	// Enum: [sha256]
	Algorithm *string `json:"algorithm"`

	// The hash value of the payload
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this intoto v002 schema content payload hash
func (m *IntotoV002SchemaContentPayloadHash) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var intotoV002SchemaContentPayloadHashTypeAlgorithmPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sha256"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		intotoV002SchemaContentPayloadHashTypeAlgorithmPropEnum = append(intotoV002SchemaContentPayloadHashTypeAlgorithmPropEnum, v)
	}
}

const (

	// IntotoV002SchemaContentPayloadHashAlgorithmSha256 captures enum value "sha256"
	IntotoV002SchemaContentPayloadHashAlgorithmSha256 string = "sha256"
)

// prop value enum
func (m *IntotoV002SchemaContentPayloadHash) validateAlgorithmEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, intotoV002SchemaContentPayloadHashTypeAlgorithmPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IntotoV002SchemaContentPayloadHash) validateAlgorithm(formats strfmt.Registry) error {

	if err := validate.Required("content"+"."+"payloadHash"+"."+"algorithm", "body", m.Algorithm); err != nil {
		return err
	}

	// value enum
	if err := m.validateAlgorithmEnum("content"+"."+"payloadHash"+"."+"algorithm", "body", *m.Algorithm); err != nil {
		return err
	}

	return nil
}

func (m *IntotoV002SchemaContentPayloadHash) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("content"+"."+"payloadHash"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this intoto v002 schema content payload hash based on the context it is used
func (m *IntotoV002SchemaContentPayloadHash) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *IntotoV002SchemaContentPayloadHash) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntotoV002SchemaContentPayloadHash) UnmarshalBinary(b []byte) error {
	var res IntotoV002SchemaContentPayloadHash
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
