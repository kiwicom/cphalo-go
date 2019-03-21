package cphalo

import (
	"encoding/json"
	"strings"
)

// IPList is a list of strings, which marshals into string and vice versa.
type IPList []string

// MarshalJSON is used by marshaler interface.
func (i IPList) MarshalJSON() ([]byte, error) {
	return json.Marshal(strings.Join(i, ","))
}

// UnmarshalJSON is used by unmarshaler interface.
func (i *IPList) UnmarshalJSON(in []byte) error {
	var replacer = strings.NewReplacer(
		`"`, "",
		`\r\n`, "",
		`\r`, "",
		`\n`, "",
	)

	var output []string

	for _, s := range strings.Split(replacer.Replace(string(in)), ",") {
		output = append(output, strings.TrimSpace(s))
	}

	*i = output

	return nil
}
