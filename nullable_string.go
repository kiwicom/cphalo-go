package cphalo

import (
	"bytes"
	"strings"
)

// NullableString is a string, which marshalls empty string into null and vice versa.
type NullableString string

func (s NullableString) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	if len(string(s)) == 0 {
		buf.WriteString(`null`)
	} else {
		buf.WriteString(`"` + string(s) + `"`)
	}

	return buf.Bytes(), nil
}

func (s *NullableString) UnmarshalJSON(in []byte) error {
	str := string(in)
	if str == `null` {
		*s = ""
		return nil
	}

	*s = NullableString(strings.Trim(str, `"`))

	return nil
}
