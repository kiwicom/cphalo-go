package cphalo

import (
	"bytes"
	"strings"
)

// NullableString is a string, which marshals empty string into null and vice versa.
type NullableString string

// MarshalJSON is used by marshaler interface.
func (s NullableString) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	if len(string(s)) == 0 {
		buf.WriteString(`null`)
	} else {
		buf.WriteString(`"` + string(s) + `"`)
	}

	return buf.Bytes(), nil
}

// UnmarshalJSON is used by unmarshaler interface.
func (s *NullableString) UnmarshalJSON(in []byte) error {
	str := string(in)
	if str == `null` {
		*s = ""
		return nil
	}

	*s = NullableString(strings.Trim(str, `"`))

	return nil
}
