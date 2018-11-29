package api

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestClient_ListCSPAccounts(t *testing.T) {
	var err error
	expectedResults := 1
	expectedID := "920b3f30-9204-469a-967c-878aa4a77c06"

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "csp_accounts_list", http.StatusOK),
			t,
			http.MethodGet,
			"/v1/csp_accounts",
			nil,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.BaseUrl, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	resp, err := client.ListCSPAccounts()

	if err != nil {
		t.Fatalf("CSP accounts list failed: %v", err)
	}

	if resp.Count != expectedResults {
		t.Errorf("expected count to be %d; got %d", expectedResults, resp.Count)
	}

	if len(resp.CSPAccounts) != expectedResults {
		t.Errorf("expected %d CSP accounts; got %d", expectedResults, resp.Count)
	}

	if resp.CSPAccounts[0].ID != expectedID {
		t.Errorf("expected CSP account 0 to have ID %s; got %s", expectedID, resp.CSPAccounts[0].ID)
	}
}

func TestClient_CreateCSPAccount(t *testing.T) {
	var err error
	reqBody := CreateCSPAccountRequest{}

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "csp_accounts_create", http.StatusCreated),
			t,
			http.MethodPost,
			"/v1/csp_accounts",
			&reqBody,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.BaseUrl, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	expectedExternalID := "this-is-external-id-1"
	resp, err := client.CreateCSPAccount(CreateCSPAccountRequest{
		RoleArn:    "arn:aws:iam::1234567890:role/CloudPassage-Service-Role",
		GroupID:    "fff04606e97b11e111d9252f8ed31222",
		ExternalID: expectedExternalID,
	})

	if err != nil {
		t.Fatalf("CSP account deletion failed: %v", err)
	}

	expectedID := "some-created-id"
	if resp != CreateCSPAccountResponse(expectedID) {
		t.Errorf("expected response to containt ID=%s; got %s", expectedID, resp)
	}

	if reqBody.CSPAccountType != "AWS" {
		t.Errorf("expected CSPAccountType to be AWS; got %s", reqBody.CSPAccountType)
	}

	if reqBody.ExternalID != expectedExternalID {
		t.Errorf("expected ExternalID to be %s; got %s", expectedExternalID, reqBody.ExternalID)
	}
}

func TestClient_DeleteCSPAccount(t *testing.T) {
	var err error

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "", http.StatusNoContent),
			t,
			http.MethodDelete,
			"/v1/csp_accounts/test",
			nil,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.BaseUrl, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	err = client.DeleteCSPAccount("test")

	if err != nil {
		t.Fatalf("CSP account deletion failed: %v", err)
	}
}
