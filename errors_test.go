package cphalo

import (
	"reflect"
	"testing"
)

func TestResponseError(t *testing.T) {
	tests := []struct {
		err  ResponseError
		code int
		msg  string
	}{
		{
			ResponseErrorGeneral{"test", 404},
			404,
			"test",
		},
		{
			ResponseError400{"test", 400},
			400,
			"request failed with 400: test",
		},
		{
			ResponseError404{"resource", "field", "value"},
			404,
			"resource resource with field=value not found",
		},
		{
			ResponseError422{Message: "message", Errors: []struct {
				Field   string `json:"field"`
				Value   string `json:"value"`
				Code    string `json:"code"`
				Details string `json:"details"`
			}{
				{"field", "value", "code", "details"},
			}},
			422,
			"validation failed for 1 fields with msg: message\n- field code",
		},
		{
			ResponseError429{},
			429,
			"Too Many Requests",
		},
		{
			ResponseError500{"message", 500},
			500,
			"server failed with code 500 and expectedErr: message",
		},
	}

	for _, tt := range tests {
		t.Run(reflect.TypeOf(tt.err).String(), func(t *testing.T) {
			if tt.code != tt.err.GetStatusCode() {
				t.Errorf("expected code %d; got %d", tt.code, tt.err.GetStatusCode())
			}
			if tt.msg != tt.err.Error() {
				t.Errorf("expected msg %q; got %q", tt.msg, tt.err.Error())
			}
		})
	}
}
