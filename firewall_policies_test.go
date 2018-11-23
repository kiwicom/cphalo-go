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
