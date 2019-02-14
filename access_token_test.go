package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestClient_RenewAccessToken(t *testing.T) {
	tests := []struct {
		name          string
		handler       func(w http.ResponseWriter, r *http.Request)
		expectedError string
	}{
		{
			"valid_access_token",
			func(w http.ResponseWriter, r *http.Request) {
				b, err := ioutil.ReadFile("testdata/access_token.json")

				if err != nil {
					t.Fatalf("cannot read file: %v", err)
				}

				fmt.Fprint(w, string(b))
			},
			"",
		},
		{
			"invalid_access_token",
			func(w http.ResponseWriter, r *http.Request) {
				b, err := ioutil.ReadFile("testdata/error_invalid_token.json")

				if err != nil {
					t.Fatalf("cannot read file: %v", err)
				}

				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, string(b))
			},
			"invalid credentials",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			ts := httptest.NewServer(http.HandlerFunc(tt.handler))
			defer ts.Close()

			client := NewClient("", "")
			client.baseURL, err = url.Parse(ts.URL)

			if err != nil {
				t.Fatalf("cannot parse test url: %v", err)
			}

			err = client.renewAccessToken()

			if len(tt.expectedError) > 0 {
				if tt.expectedError != err.Error() {
					t.Errorf("expected error: %s; got %s", tt.expectedError, err.Error())
				}
				return
			}

			if err != nil {
				t.Fatalf("renewal failed: %v", err)
			}

			expectedToken := "some_token_for_cp_halo_rest_api1"
			if client.accessToken != expectedToken {
				t.Errorf("Client's access token not properly set, expected %s; got %s", expectedToken, client.accessToken)
			}
		})
	}
}
