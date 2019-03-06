package cphalo

import (
	"bytes"
	"fmt"
)

// StringableBool is a bool, which marshals bool into string and vice versa.
type StringableBool bool

// MarshalJSON is used by marshaler interface.
func (b StringableBool) MarshalJSON() ([]byte, error) {
	fmt.Printf("RAW MARSHALL: %#v\n", b)
	var buf bytes.Buffer
	if b {
		buf.WriteString("true")
	} else {
		buf.WriteString("false")
	}

	return buf.Bytes(), nil
}

// UnmarshalJSON is used by unmarshaler interface.
func (b *StringableBool) UnmarshalJSON(in []byte) error {
	*b = string(in) == "true"

	return nil
}
