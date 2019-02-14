package api

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestClient_ListFirewallInterfaces(t *testing.T) {
	var err error
	expectedResults := 7
	expectedID := "eab260c09956012ee2db4087123ad87s"

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "firewall_interfaces_list", http.StatusOK),
			t,
			http.MethodGet,
			"/v1/firewall_interfaces",
			nil,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.baseUrl, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	resp, err := client.ListFirewallInterfaces()

	if err != nil {
		t.Fatalf("firewall policies list failed: %v", err)
	}

	if resp.Count != expectedResults {
		t.Errorf("expected count to be %d; got %d", expectedResults, resp.Count)
	}

	if len(resp.Interfaces) != expectedResults {
		t.Errorf("expected %d firewall policies; got %d", expectedResults, resp.Count)
	}

	if resp.Interfaces[0].ID != expectedID {
		t.Errorf("expected firewall interface 0 to have ID %s; got %s", expectedID, resp.Interfaces[0].ID)
	}
}

func TestClient_GetFirewallInterface(t *testing.T) {
	var err error
	expectedID := "eab41be09956012ee2db4087123ad87s"

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "firewall_interfaces_get", http.StatusOK),
			t,
			http.MethodGet,
			"/v1/firewall_interfaces/id",
			nil,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.baseUrl, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	resp, err := client.GetFirewallInterface("id")

	if err != nil {
		t.Fatalf("firewall policies get failed: %v", err)
	}

	if resp.Interface.ID != expectedID {
		t.Errorf("expected firewall interface to have ID %s; got %s", expectedID, resp.Interface.ID)
	}
}

func TestClient_CreateFirewallInterface(t *testing.T) {
	var err error
	reqBody := CreateFirewallInterfaceRequest{}

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "firewall_interfaces_get", http.StatusOK),
			t,
			http.MethodPost,
			"/v1/firewall_interfaces",
			&reqBody,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.baseUrl, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	fwInterface := FirewallInterface{
		ID:   "id",
		Name: "hello",
	}

	resp, err := client.CreateFirewallInterface(fwInterface)

	if err != nil {
		t.Fatalf("Firewall interface creating failed: %v", err)
	}

	expectedID := "eab41be09956012ee2db4087123ad87s"
	if resp.Interface.ID != expectedID {
		t.Errorf("expected response to containt ID=%s; got %s", expectedID, resp.Interface.ID)
	}

	if reqBody.Interface.ID != fwInterface.ID {
		t.Errorf("expected request to contain ID=%s; got %s", fwInterface.ID, reqBody.Interface.ID)
	}

	if reqBody.Interface.Name != fwInterface.Name {
		t.Errorf("expected request to contain Name=%s; got %s", fwInterface.Name, reqBody.Interface.Name)
	}
}

func TestClient_UpdateFirewallInterface(t *testing.T) {
	var err error
	reqBody := UpdateFirewallInterfaceRequest{}

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "", http.StatusOK),
			t,
			http.MethodPut,
			"/v1/firewall_interfaces/id",
			&reqBody,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.baseUrl, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	fwInterface := FirewallInterface{
		ID:   "id",
		Name: "hello",
	}

	err = client.UpdateFirewallInterface(fwInterface)

	if err != nil {
		t.Fatalf("Firewall interface updating failed: %v", err)
	}

	if reqBody.Interface.ID != fwInterface.ID {
		t.Errorf("expected request to contain ID=%s; got %s", fwInterface.ID, reqBody.Interface.ID)
	}

	if reqBody.Interface.Name != fwInterface.Name {
		t.Errorf("expected request to contain Name=%s; got %s", fwInterface.Name, reqBody.Interface.Name)
	}
}

func TestClient_DeleteFirewallInterface(t *testing.T) {
	var err error

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "", http.StatusNoContent),
			t,
			http.MethodDelete,
			"/v1/firewall_interfaces/id",
			nil,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.baseUrl, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	err = client.DeleteFirewallInterface("id")

	if err != nil {
		t.Fatalf("Firewall interface deletion failed: %v", err)
	}
}
