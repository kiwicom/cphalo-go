package cphalo

import (
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
	client.baseURL, err = url.Parse(ts.URL)

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
	client.baseURL, err = url.Parse(ts.URL)

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

func TestClient_CreateFirewallRule(t *testing.T) {
	var err error
	reqBody := CreateFirewallRuleRequest{}

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "firewall_rules_get", http.StatusOK),
			t,
			http.MethodPost,
			"/v1/firewall_policies/123/firewall_rules",
			&reqBody,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.baseURL, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	rule := FirewallRule{
		ID:               "id",
		Chain:            "OUTPUT",
		Action:           "DROP",
		Position:         1,
		ConnectionStates: "NEW",
		Log:              true,
		LogPrefix:        "test_",
		Comment:          "test",
	}

	resp, err := client.CreateFirewallRule("123", rule)

	if err != nil {
		t.Fatalf("Firewall rule updating failed: %v", err)
	}

	expectedID := "jtap810cee5b11e89d0d15ea2eff2521"
	if resp.Rule.ID != expectedID {
		t.Errorf("expected response to containt ID=%s; got %s", expectedID, resp.Rule.ID)
	}

	if reqBody.Rule.ID != rule.ID {
		t.Errorf("expected request to contain ID=%s; got %s", rule.ID, reqBody.Rule.ID)
	}

	if reqBody.Rule.Chain != rule.Chain {
		t.Errorf("expected request to contain Chain=%s; got %s", rule.Chain, reqBody.Rule.Chain)
	}

	if reqBody.Rule.Action != rule.Action {
		t.Errorf("expected request to contain Action=%s; got %s", rule.Action, reqBody.Rule.Action)
	}

	if reqBody.Rule.Position != rule.Position {
		t.Errorf("expected request to contain Position=%d; got %d", rule.Position, reqBody.Rule.Position)
	}

	if reqBody.Rule.ConnectionStates != rule.ConnectionStates {
		t.Errorf("expected request to contain ConnectionStates=%s; got %s", rule.ConnectionStates, reqBody.Rule.ConnectionStates)
	}

	if reqBody.Rule.Log != rule.Log {
		t.Errorf("expected request to contain Log=%t; got %t", rule.Log, reqBody.Rule.Log)
	}

	if reqBody.Rule.LogPrefix != rule.LogPrefix {
		t.Errorf("expected request to contain LogPrefix=%s; got %s", rule.LogPrefix, reqBody.Rule.LogPrefix)
	}

	if reqBody.Rule.Comment != rule.Comment {
		t.Errorf("expected request to contain Comment=%s; got %s", rule.Comment, reqBody.Rule.Comment)
	}
}

func TestClient_UpdateFirewallRule(t *testing.T) {
	var err error
	reqBody := UpdateFirewallRuleRequest{}

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "", http.StatusOK),
			t,
			http.MethodPut,
			"/v1/firewall_policies/123/firewall_rules/id",
			&reqBody,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.baseURL, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	rule := FirewallRule{
		ID:               "id",
		Chain:            "OUTPUT",
		Action:           "DROP",
		Position:         1,
		ConnectionStates: "NEW",
		Log:              true,
		LogPrefix:        "test_",
		Comment:          "test",
	}

	err = client.UpdateFirewallRule("123", rule)

	if err != nil {
		t.Fatalf("Firewall rule updating failed: %v", err)
	}

	if reqBody.Rule.ID != rule.ID {
		t.Errorf("expected response to contain ID=%s; got %s", rule.ID, reqBody.Rule.ID)
	}

	if reqBody.Rule.Chain != rule.Chain {
		t.Errorf("expected response to contain Chain=%s; got %s", rule.Chain, reqBody.Rule.Chain)
	}

	if reqBody.Rule.Action != rule.Action {
		t.Errorf("expected response to contain Action=%s; got %s", rule.Action, reqBody.Rule.Action)
	}

	if reqBody.Rule.Position != rule.Position {
		t.Errorf("expected response to contain Position=%d; got %d", rule.Position, reqBody.Rule.Position)
	}

	if reqBody.Rule.ConnectionStates != rule.ConnectionStates {
		t.Errorf("expected response to contain ConnectionStates=%s; got %s", rule.ConnectionStates, reqBody.Rule.ConnectionStates)
	}

	if reqBody.Rule.Log != rule.Log {
		t.Errorf("expected response to contain Log=%t; got %t", rule.Log, reqBody.Rule.Log)
	}

	if reqBody.Rule.LogPrefix != rule.LogPrefix {
		t.Errorf("expected response to contain LogPrefix=%s; got %s", rule.LogPrefix, reqBody.Rule.LogPrefix)
	}

	if reqBody.Rule.Comment != rule.Comment {
		t.Errorf("expected response to contain Comment=%s; got %s", rule.Comment, reqBody.Rule.Comment)
	}
}

func TestClient_DeleteFirewallRule(t *testing.T) {
	var err error

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "", http.StatusNoContent),
			t,
			http.MethodDelete,
			"/v1/firewall_policies/123/firewall_rules/id",
			nil,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.baseURL, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	err = client.DeleteFirewallRule("123", "id")

	if err != nil {
		t.Fatalf("Firewall rule deletion failed: %v", err)
	}
}
