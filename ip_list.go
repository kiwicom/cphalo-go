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

	output := []string{}

	for _, s := range strings.Split(replacer.Replace(string(in)), ",") {
		trimmed := strings.TrimSpace(s)
		if len(trimmed) > 0 {
			output = append(output, trimmed)
		}
	}

	*i = output

	return nil
}
