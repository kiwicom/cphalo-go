package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestClient_ListFirewallRules(t *testing.T) {
	var err error
	expectedResults := 2
	expectedID := "ce3d810cee5b23ju9d0d15ea2eff2521"

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "firewall_rules_list", http.StatusOK),
			t,
			http.MethodGet,
			"/v1/firewall_policies/123/firewall_rules",
			nil,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.BaseUrl, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	resp, err := client.ListFirewallRules("123")

	if err != nil {
		t.Fatalf("firewall rules list failed: %v", err)
	}

	if resp.Count != expectedResults {
		t.Errorf("expected count to be %d; got %d", expectedResults, resp.Count)
	}

	fmt.Println(resp)

	if len(resp.Rules) != expectedResults {
		t.Errorf("expected %d firewall rules; got %d", expectedResults, resp.Count)
	}

	if resp.Rules[0].ID != expectedID {
		t.Errorf("expected firewall rule 0 to have ID %s; got %s", expectedID, resp.Rules[0].ID)
	}
}

func TestClient_GetFirewallRule(t *testing.T) {
	var err error
	expectedID := "jtap810cee5b11e89d0d15ea2eff2521"

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "firewall_rules_get", http.StatusOK),
			t,
			http.MethodGet,
			"/v1/firewall_policies/123/firewall_rules/id",
			nil,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.BaseUrl, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	resp, err := client.GetFirewallRule("123", "id")

	if err != nil {
		t.Fatalf("firewall rules get failed: %v", err)
	}

	if resp.Rule.ID != expectedID {
		t.Errorf("expected firewall rule to have ID %s; got %s", expectedID, resp.Rule.ID)
	}
}
