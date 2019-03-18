package cphalo

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestClient_GetServerGroupFirewallPolicy(t *testing.T) {
	var (
		err        error
		expectedID = NullableString("8eb1b050abe2013295e406ba9a9c633c")
		groupID    = "random-group-id"
	)

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "server_groups_get", http.StatusOK),
			t,
			http.MethodGet,
			fmt.Sprintf("/v1/groups/%s", groupID),
			nil,
		),
	)
	defer ts.Close()

	client := NewClient("", "", nil)
	client.baseURL, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	resp, err := client.GetServerGroupFirewallPolicy(groupID)

	if err != nil {
		t.Fatalf("server groups get failed: %v", err)
	}

	if resp.Group.LinuxFirewallPolicyID != expectedID {
		t.Errorf("expected group to have ID %s; got %s", expectedID, resp.Group.LinuxFirewallPolicyID)
	}
}

func TestClient_UpdateServerGroupFirewallPolicy(t *testing.T) {
	var (
		err      error
		reqBody  = UpdateServerGroupFirewallPolicyRequest{}
		policyID = NullableString("random-policy-id")
		groupID  = "group-id"
	)

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "server_groups_get", http.StatusCreated),
			t,
			http.MethodPut,
			fmt.Sprintf("/v1/groups/%s", groupID),
			&reqBody,
		),
	)
	defer ts.Close()

	client := NewClient("", "", nil)
	client.baseURL, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	err = client.UpdateServerGroupFirewallPolicy(ServerGroupFirewallPolicy{
		GroupID:               groupID,
		LinuxFirewallPolicyID: policyID,
	})

	if err != nil {
		t.Fatalf("CSP account creating failed: %v", err)
	}

	if reqBody.Group.LinuxFirewallPolicyID != policyID {
		t.Errorf("expected request to have linux firewall policy id %s; got %s", policyID, reqBody.Group.LinuxFirewallPolicyID)
	}
}
