package cphalo

import (
	"encoding/json"
	"testing"
)

var (
	tests []struct {
		name string
		ns   NullableString
		s    string
	}
)

func init() {
	tests = []struct {
		name string
		ns   NullableString
		s    string
	}{
		{"simple", NullableString("test"), `"test"`},
		{"space", NullableString(" "), `" "`},
		{"empty", NullableString(""), `null`},
	}
}

func TestNullableString_MarshalJSON(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := json.Marshal(tt.ns)

			if err != nil {
				t.Fatalf("cannot marshall: %v", err)
			}

			got := string(b)
			if got != tt.s {
				t.Errorf("s %s; got %s", tt.s, string(got))
			}
		})
	}
}

func TestNullableString_UnmarshalJSON(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NullableString("")

			if err := json.Unmarshal([]byte(tt.s), &got); err != nil {
				t.Fatalf("unmarshalling failed: %v", err)
			}

			if got != tt.ns {
				t.Errorf("expected %s; got %s", tt.ns, string(got))
			}
		})
	}
}
