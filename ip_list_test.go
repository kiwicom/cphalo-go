package cphalo

import (
	"encoding/json"
	"reflect"
	"testing"
)

type IPListTestItem struct {
	name   string
	ips    IPList
	strVal string
}

func TestIPList_MarshalJSON(t *testing.T) {
	tests := []IPListTestItem{
		{"single", IPList{"0.0.0.0/0"}, `"0.0.0.0/0"`},
		{"double", IPList{"1.1.1.1", "2.2.2.2"}, `"1.1.1.1,2.2.2.2"`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := json.Marshal(tt.ips)

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

func TestIPList_UnmarshalJSON(t *testing.T) {
	tests := []IPListTestItem{
		{"single", IPList{"0.0.0.0/0"}, `"0.0.0.0/0"`},
		{"double", IPList{"1.1.1.1", "2.2.2.2"}, `"1.1.1.1,2.2.2.2"`},
		{"newlines", IPList{"1.1.1.1", "2.2.2.2"}, `"1.1.1.1,\r\n2.2.2.2"`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IPList{}

			if err := json.Unmarshal([]byte(tt.strVal), &got); err != nil {
				t.Fatalf("unmarshalling failed: %v", err)
			}

			if !reflect.DeepEqual(got, tt.ips) {
				t.Fatalf("expected %s; got %s", tt.ips, got)
			}
		})
	}
}
