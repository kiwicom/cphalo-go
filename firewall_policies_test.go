package api

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestClient_ListFirewallPolicies(t *testing.T) {
	var err error
	expectedResults := 3
	expectedID := "8cb238a0ee5511e18s9a4d1cedf20253"

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "firewall_policies_list", http.StatusOK),
			t,
			http.MethodGet,
			"/v1/firewall_policies",
			nil,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.BaseUrl, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	resp, err := client.ListFirewallPolicies()

	if err != nil {
		t.Fatalf("firewall policies list failed: %v", err)
	}

	if resp.Count != expectedResults {
		t.Errorf("expected count to be %d; got %d", expectedResults, resp.Count)
	}

	if len(resp.Policies) != expectedResults {
		t.Errorf("expected %d firewall policies; got %d", expectedResults, resp.Count)
	}

	if resp.Policies[0].ID != expectedID {
		t.Errorf("expected firewall policy 0 to have ID %s; got %s", expectedID, resp.Policies[0].ID)
	}
}

func TestClient_GetFirewallPolicy(t *testing.T) {
	var err error
	expectedID := "be28b106ee5b11e8a7s1017da54e9117"

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "firewall_policies_get", http.StatusOK),
			t,
			http.MethodGet,
			"/v1/firewall_policies/id",
			nil,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.BaseUrl, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	resp, err := client.GetFirewallPolicy("id")

	if err != nil {
		t.Fatalf("firewall policies get failed: %v", err)
	}

	if resp.Policy.ID != expectedID {
		t.Errorf("expected firewall policy to have ID %s; got %s", expectedID, resp.Policy.ID)
	}
}

func TestClient_CreateFirewallPolicy(t *testing.T) {
	var err error
	reqBody := CreateFirewallPolicyRequest{}

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "firewall_policies_get", http.StatusOK),
			t,
			http.MethodPost,
			"/v1/firewall_policies",
			&reqBody,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.BaseUrl, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	policy := FirewallPolicy{
		ID:                    "id",
		Platform:              "linux",
		Name:                  "hello",
		Shared:                true,
		IgnoreForwardingRules: false,
		Description:           "hai",
	}

	resp, err := client.CreateFirewallPolicy(policy)

	if err != nil {
		t.Fatalf("Firewall policy creating failed: %v", err)
	}

	expectedID := "be28b106ee5b11e8a7s1017da54e9117"
	if resp.Policy.ID != expectedID {
		t.Errorf("expected response to containt ID=%s; got %s", expectedID, resp.Policy.ID)
	}

	if reqBody.Policy.ID != policy.ID {
		t.Errorf("expected request to contain ID=%s; got %s", policy.ID, reqBody.Policy.ID)
	}

	if reqBody.Policy.Platform != policy.Platform {
		t.Errorf("expected request to contain Platform=%s; got %s", policy.Platform, reqBody.Policy.Platform)
	}

	if reqBody.Policy.Name != policy.Name {
		t.Errorf("expected request to contain Name=%s; got %s", policy.Name, reqBody.Policy.Name)
	}

	if reqBody.Policy.Shared != policy.Shared {
		t.Errorf("expected request to contain Shared=%t; got %t", policy.Shared, reqBody.Policy.Shared)
	}

	if reqBody.Policy.IgnoreForwardingRules != policy.IgnoreForwardingRules {
		t.Errorf("expected request to contain IgnoreForwardingRules=%t; got %t", policy.IgnoreForwardingRules, reqBody.Policy.IgnoreForwardingRules)
	}

	if reqBody.Policy.Description != policy.Description {
		t.Errorf("expected request to contain Description=%s; got %s", policy.Description, reqBody.Policy.Description)
	}
}

func TestClient_UpdateFirewallPolicy(t *testing.T) {
	var err error
	reqBody := UpdateFirewallPolicyRequest{}

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "", http.StatusOK),
			t,
			http.MethodPut,
			"/v1/firewall_policies/id",
			&reqBody,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.BaseUrl, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	policy := FirewallPolicy{
		ID:                    "id",
		Platform:              "linux",
		Name:                  "hello",
		Shared:                true,
		IgnoreForwardingRules: false,
		Description:           "hai",
	}

	err = client.UpdateFirewallPolicy(policy)

	if err != nil {
		t.Fatalf("Firewall policy updating failed: %v", err)
	}

	if reqBody.Policy.ID != policy.ID {
		t.Errorf("expected request to contain ID=%s; got %s", policy.ID, reqBody.Policy.ID)
	}

	if reqBody.Policy.Platform != policy.Platform {
		t.Errorf("expected request to contain Platform=%s; got %s", policy.Platform, reqBody.Policy.Platform)
	}

	if reqBody.Policy.Name != policy.Name {
		t.Errorf("expected request to contain Name=%s; got %s", policy.Name, reqBody.Policy.Name)
	}

	if reqBody.Policy.Shared != policy.Shared {
		t.Errorf("expected request to contain Shared=%t; got %t", policy.Shared, reqBody.Policy.Shared)
	}

	if reqBody.Policy.IgnoreForwardingRules != policy.IgnoreForwardingRules {
		t.Errorf("expected request to contain IgnoreForwardingRules=%t; got %t", policy.IgnoreForwardingRules, reqBody.Policy.IgnoreForwardingRules)
	}

	if reqBody.Policy.Description != policy.Description {
		t.Errorf("expected request to contain Description=%s; got %s", policy.Description, reqBody.Policy.Description)
	}
}

func TestClient_DeleteFirewallPolicy(t *testing.T) {
	var err error

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "", http.StatusNoContent),
			t,
			http.MethodDelete,
			"/v1/firewall_policies/id",
			nil,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.BaseUrl, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	err = client.DeleteFirewallPolicy("id")

	if err != nil {
		t.Fatalf("Firewall policy deletion failed: %v", err)
	}
}
