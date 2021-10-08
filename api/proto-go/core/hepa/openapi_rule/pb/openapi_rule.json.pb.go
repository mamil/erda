// Code generated by protoc-gen-go-json. DO NOT EDIT.
// Source: openapi_rule.proto

package pb

import (
	bytes "bytes"
	json "encoding/json"
	jsonpb "github.com/erda-project/erda-infra/pkg/transport/http/encoding/jsonpb"
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "encoding/json" package it is being compiled against.
var _ json.Marshaler = (*DeleteLimitResponse)(nil)
var _ json.Unmarshaler = (*DeleteLimitResponse)(nil)
var _ json.Marshaler = (*DeleteLimitRequest)(nil)
var _ json.Unmarshaler = (*DeleteLimitRequest)(nil)
var _ json.Marshaler = (*LimitRuleInfo)(nil)
var _ json.Unmarshaler = (*LimitRuleInfo)(nil)
var _ json.Marshaler = (*UpdateLimitResponse)(nil)
var _ json.Unmarshaler = (*UpdateLimitResponse)(nil)
var _ json.Marshaler = (*UpdateLimitRequest)(nil)
var _ json.Unmarshaler = (*UpdateLimitRequest)(nil)
var _ json.Marshaler = (*CreateLimitResponse)(nil)
var _ json.Unmarshaler = (*CreateLimitResponse)(nil)
var _ json.Marshaler = (*LimitType)(nil)
var _ json.Unmarshaler = (*LimitType)(nil)
var _ json.Marshaler = (*LimitRequest)(nil)
var _ json.Unmarshaler = (*LimitRequest)(nil)
var _ json.Marshaler = (*CreateLimitRequest)(nil)
var _ json.Unmarshaler = (*CreateLimitRequest)(nil)
var _ json.Marshaler = (*GetLimitsRequest)(nil)
var _ json.Unmarshaler = (*GetLimitsRequest)(nil)
var _ json.Marshaler = (*GetLimitsResponse)(nil)
var _ json.Unmarshaler = (*GetLimitsResponse)(nil)

// DeleteLimitResponse implement json.Marshaler.
func (m *DeleteLimitResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteLimitResponse implement json.Marshaler.
func (m *DeleteLimitResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// DeleteLimitRequest implement json.Marshaler.
func (m *DeleteLimitRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// DeleteLimitRequest implement json.Marshaler.
func (m *DeleteLimitRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// LimitRuleInfo implement json.Marshaler.
func (m *LimitRuleInfo) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// LimitRuleInfo implement json.Marshaler.
func (m *LimitRuleInfo) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateLimitResponse implement json.Marshaler.
func (m *UpdateLimitResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateLimitResponse implement json.Marshaler.
func (m *UpdateLimitResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// UpdateLimitRequest implement json.Marshaler.
func (m *UpdateLimitRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// UpdateLimitRequest implement json.Marshaler.
func (m *UpdateLimitRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateLimitResponse implement json.Marshaler.
func (m *CreateLimitResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateLimitResponse implement json.Marshaler.
func (m *CreateLimitResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// LimitType implement json.Marshaler.
func (m *LimitType) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// LimitType implement json.Marshaler.
func (m *LimitType) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// LimitRequest implement json.Marshaler.
func (m *LimitRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// LimitRequest implement json.Marshaler.
func (m *LimitRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// CreateLimitRequest implement json.Marshaler.
func (m *CreateLimitRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// CreateLimitRequest implement json.Marshaler.
func (m *CreateLimitRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetLimitsRequest implement json.Marshaler.
func (m *GetLimitsRequest) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetLimitsRequest implement json.Marshaler.
func (m *GetLimitsRequest) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}

// GetLimitsResponse implement json.Marshaler.
func (m *GetLimitsResponse) MarshalJSON() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := (&jsonpb.Marshaler{
		OrigName:     false,
		EnumsAsInts:  false,
		EmitDefaults: true,
	}).Marshal(buf, m)
	return buf.Bytes(), err
}

// GetLimitsResponse implement json.Marshaler.
func (m *GetLimitsResponse) UnmarshalJSON(b []byte) error {
	return (&protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}).Unmarshal(b, m)
}