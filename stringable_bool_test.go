package cphalo

import (
	"encoding/json"
	"testing"
)

type stringableBoolTestItem struct {
	name    string
	b       StringableBool
	strVal  string
	boolVal bool
}

func getStringableBoolTestData() []stringableBoolTestItem {
	return []stringableBoolTestItem{
		{"true", StringableBool(true), `"true"`, true},
		{"false", StringableBool(false), `"false"`, false},
		{"invalid", StringableBool(false), `"false"`, false},
	}
}

func TestStringableBool_MarshalJSON(t *testing.T) {
	tests := getStringableBoolTestData()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := json.Marshal(tt.b)

			if err != nil {
				t.Fatalf("cannot marshall: %v", err)
			}

			got := string(b)
			if got != tt.strVal {
				t.Errorf("s %s; got %s", tt.strVal, got)
			}
		})
	}
}

func TestStringableBool_UnmarshalJSON(t *testing.T) {
	tests := getStringableBoolTestData()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StringableBool(true)

			if err := json.Unmarshal([]byte(tt.strVal), &got); err != nil {
				t.Fatalf("unmarshalling failed: %v", err)
			}

			if got != StringableBool(tt.boolVal) {
				t.Errorf("expected %t; got %t", tt.boolVal, got)
			}
		})
	}
}
