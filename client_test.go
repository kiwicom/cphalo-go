package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestNewClient(t *testing.T) {
	appKey := "123"
	appSecret := "321"
	client := NewClient(appKey, appSecret)

	if client == nil {
		t.Fatal("new client returned nil")
	}

	if client.AppKey != appKey {
		t.Errorf("expected app key %s; got %s", appKey, client.AppKey)
	}
	if client.AppSecret != appSecret {
		t.Errorf("expected app secret %s; got %s", appSecret, client.AppSecret)
	}
	if client.BaseUrl.String() != DefaultBaseUrl {
		t.Errorf("expected base url %s; got %s", DefaultBaseUrl, client.BaseUrl.String())
	}
	if client.Timeout != DefaultTimeout {
		t.Errorf("expected timeout %s; got %s", DefaultTimeout, client.Timeout)
	}
	if client.MaxAuthTries != DefaultMaxAuthTries {
		t.Errorf("expected max auth tries %d; got %d", DefaultMaxAuthTries, client.MaxAuthTries)
	}
	if client.client == nil {
		t.Fatal("http client not set")
	}
	if client.client.Timeout != DefaultTimeout {
		t.Errorf("expected client timeout %s; got %s", DefaultTimeout, client.client.Timeout)
	}
}

func authTestHandler(next http.Handler, t *testing.T) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.RequestURI, "/oauth") {
			b, err := ioutil.ReadFile("example_responses/access_token.json")

			if err != nil {
				t.Fatalf("cannot read file: %v", err)
			}

			w.WriteHeader(http.StatusOK)
			_, err = w.Write(b)

			if err != nil {
				t.Fatalf("cannot write response: %v", err)
			}
			return
		}
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func jsonResponseTestHandler(t *testing.T, responseFile string, code int) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)

		if responseFile == "" {
			return
		}

		b, err := ioutil.ReadFile(fmt.Sprintf("example_responses/%s.json", responseFile))

		if err != nil {
			t.Fatalf("cannot read file: %v", err)
		}

		fmt.Fprint(w, string(b))
	}

	return http.HandlerFunc(fn)
}

func requestValidatorTestHandler(next http.Handler, t *testing.T, expectedMethod, expectedURI string, body interface{}) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if t == nil {
			t.Fatal("request not set for assertion")
		}
		if r.Method != expectedMethod {
			t.Errorf("invalid method, expected %s; got %s", expectedMethod, r.Method)
		}
		if r.RequestURI != expectedURI {
			t.Errorf("invalid URI, expected %s; got %s", expectedURI, r.RequestURI)
		}

		if body != nil {
			b, err := ioutil.ReadAll(r.Body)

			if err != nil {
				t.Fatalf("reading body failed: %v", err)
			}

			if err := json.Unmarshal(b, body); err != nil {
				t.Fatalf("unmarshalling body failed: %v", err)
			}
		}

		if next != nil {
			next.ServeHTTP(w, r)
		}

	}

	return authTestHandler(http.HandlerFunc(fn), t)
}
