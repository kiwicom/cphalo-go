package api

import (
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
