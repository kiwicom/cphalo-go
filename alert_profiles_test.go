package cphalo

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestClient_ListAlertProfiles(t *testing.T) {
	var err error
	expectedResults := 2
	expectedID := "0226a27af95c11e5a92a471a4310f7c2"

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "alert_profiles_list", http.StatusOK),
			t,
			http.MethodGet,
			"/v1/alert_profiles",
			nil,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.baseURL, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	resp, err := client.ListAlertProfiles()

	if err != nil {
		t.Fatalf("alert profiles list failed: %v", err)
	}

	if resp.Count != expectedResults {
		t.Errorf("expected count to be %d; got %d", expectedResults, resp.Count)
	}

	if len(resp.AlertProfiles) != expectedResults {
		t.Errorf("expected %d alert profiles; got %d", expectedResults, len(resp.AlertProfiles))
	}

	if resp.AlertProfiles[0].ID != expectedID {
		t.Errorf("expected alert profile 0 to have ID %s; got %s", expectedID, resp.AlertProfiles[0].ID)
	}
}
