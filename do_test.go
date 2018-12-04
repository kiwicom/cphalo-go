package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func TestValidateResponse(t *testing.T) {
	tests := []struct {
		code        int
		bodyFile    string
		expectedErr interface{}
	}{
		{http.StatusTeapot, "error_general.json", &ResponseError400{}},
		{http.StatusNotFound, "error_404.json", &ResponseError404{}},
		{http.StatusUnprocessableEntity, "error_422.json", &ResponseError422{}},
		{http.StatusTooManyRequests, "", &ResponseError429{}},
		{http.StatusInternalServerError, "error_500.json", &ResponseError500{}},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("code_%d", tt.code), func(t *testing.T) {
			var err error
			var b []byte

			if len(tt.bodyFile) > 0 {

				path := fmt.Sprintf("example_responses/%s", tt.bodyFile)
				b, err = ioutil.ReadFile(path)

				if err != nil {
					t.Fatalf("cannot read file %s: %v", path, err)
				}
			}

			resp := http.Response{StatusCode: tt.code, Body: ioutil.NopCloser(bytes.NewBuffer(b))}
			err = validateResponse(&resp)

			actualErrType := reflect.TypeOf(err)
			expectedErrType := reflect.TypeOf(tt.expectedErr)

			if actualErrType != expectedErrType {
				t.Errorf("expected expectedErr type %s; got %s", expectedErrType, actualErrType)
			}
		})
	}
}
