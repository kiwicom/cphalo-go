package cphalo

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// StringableBool is a bool, which marshals bool into string and vice versa.
type StringableBool bool

// MarshalJSON is used by marshaler interface.
func (b StringableBool) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.FormatBool(bool(b)))
}

// UnmarshalJSON is used by unmarshaler interface.
func (b *StringableBool) UnmarshalJSON(in []byte) error {
	s, err := strconv.Unquote(string(in))

	if err == strconv.ErrSyntax {
		s = string(in)
	} else if err != nil {
		return fmt.Errorf("could not unquote string: %v", err)
	}

	*b = s == "true"

	return nil
}
