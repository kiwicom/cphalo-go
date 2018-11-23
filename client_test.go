package api

import (
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

func jsonResponseTestHandler(name string, t *testing.T, auth bool) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadFile(fmt.Sprintf("example_responses/%s.json", name))

		if err != nil {
			t.Fatalf("cannot read file: %v", err)
		}

		fmt.Fprint(w, string(b))
	}

	if auth {
		return authTestHandler(http.HandlerFunc(fn), t)
	}

	return http.HandlerFunc(fn)
}
