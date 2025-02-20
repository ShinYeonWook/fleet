// Package optjson provides types that can be used to represent optional JSON
// values. The Set field indicates if the value was set (provided in the
// unmarshaled JSON payload) while the Valid field indicates if the value was
// not null.
//
// The types also support marshaling, respecting the Valid field (it will,
// however, be marshaled even if Set is false, although it will marshal to
// null).
//
// Inspired by https://www.calhoun.io/how-to-determine-if-a-json-key-has-been-set-to-null-or-not-provided/
package optjson

import (
	"bytes"
	"encoding/json"
)

// String represents an optional string value.
type String struct {
	Set   bool
	Valid bool
	Value string
}

func SetString(s string) String {
	return String{Set: true, Valid: true, Value: s}
}

func (s String) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.Value)
}

func (s *String) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	s.Set = true
	s.Valid = false

	if bytes.Equal(data, []byte("null")) {
		// The key was set to null, blank the value
		s.Value = ""
		return nil
	}

	// The key isn't set to null
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	s.Value = v
	s.Valid = true
	return nil
}

// Bool represents an optional boolean value.
type Bool struct {
	Set   bool
	Valid bool
	Value bool
}

func SetBool(b bool) Bool {
	return Bool{Set: true, Valid: true, Value: b}
}

func (b Bool) MarshalJSON() ([]byte, error) {
	if !b.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(b.Value)
}

func (b *Bool) UnmarshalJSON(data []byte) error {
	// If this method was called, the value was set.
	b.Set = true
	b.Valid = false

	if bytes.Equal(data, []byte("null")) {
		// The key was set to null, blank the value
		b.Value = false
		return nil
	}

	// The key isn't set to null
	var v bool
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	b.Value = v
	b.Valid = true
	return nil
}
