package cphalo

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func TestClient_ListFirewallZones(t *testing.T) {
	var err error
	expectedResults := 1
	expectedID := "ea81ec609956012ee2db40989asd0980"

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "firewall_zones_list", http.StatusOK),
			t,
			http.MethodGet,
			"/v1/firewall_zones",
			nil,
		),
	)
	defer ts.Close()

	client := NewClient("", "", nil)
	client.baseURL, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	resp, err := client.ListFirewallZones()

	if err != nil {
		t.Fatalf("firewall zones list failed: %v", err)
	}

	if resp.Count != expectedResults {
		t.Errorf("expected count to be %d; got %d", expectedResults, resp.Count)
	}

	if len(resp.Zones) != expectedResults {
		t.Errorf("expected %d firewall zones; got %d", expectedResults, resp.Count)
	}

	if resp.Zones[0].ID != expectedID {
		t.Errorf("expected firewall zone 0 to have ID %s; got %s", expectedID, resp.Zones[0].ID)
	}
}

func TestClient_GetFirewallZone(t *testing.T) {
	var err error
	expectedID := "ea81ec609956012ee2db40989asd0980"

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "firewall_zones_get", http.StatusOK),
			t,
			http.MethodGet,
			"/v1/firewall_zones/id",
			nil,
		),
	)
	defer ts.Close()

	client := NewClient("", "", nil)
	client.baseURL, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	resp, err := client.GetFirewallZone("id")

	if err != nil {
		t.Fatalf("firewall zones get failed: %v", err)
	}

	if resp.Zone.ID != expectedID {
		t.Errorf("expected firewall zone to have ID %s; got %s", expectedID, resp.Zone.ID)
	}
}

func TestClient_CreateFirewallZone(t *testing.T) {
	var err error
	reqBody := CreateFirewallZoneRequest{}

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "firewall_zones_get", http.StatusOK),
			t,
			http.MethodPost,
			"/v1/firewall_zones",
			&reqBody,
		),
	)
	defer ts.Close()

	client := NewClient("", "", nil)
	client.baseURL, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	service := FirewallZone{
		ID:        "id",
		Name:      "hello",
		IPAddress: IPList{"0.0.0.0/0"},
	}

	resp, err := client.CreateFirewallZone(service)

	if err != nil {
		t.Fatalf("firewall zone creating failed: %v", err)
	}

	expectedID := "ea81ec609956012ee2db40989asd0980"
	if resp.Zone.ID != expectedID {
		t.Errorf("expected response to containt ID=%s; got %s", expectedID, resp.Zone.ID)
	}

	if reqBody.Zone.ID != service.ID {
		t.Errorf("expected request to contain ID=%s; got %s", service.ID, reqBody.Zone.ID)
	}

	if reqBody.Zone.Name != service.Name {
		t.Errorf("expected request to contain Name=%s; got %s", service.Name, reqBody.Zone.Name)
	}

	if !reflect.DeepEqual(reqBody.Zone.IPAddress, service.IPAddress) {
		t.Errorf("expected request to contain IPAddress=%s; got %s", service.IPAddress, reqBody.Zone.IPAddress)
	}
}

func TestClient_UpdateFirewallZone(t *testing.T) {
	var err error
	reqBody := UpdateFirewallZoneRequest{}

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "", http.StatusOK),
			t,
			http.MethodPut,
			"/v1/firewall_zones/id",
			&reqBody,
		),
	)
	defer ts.Close()

	client := NewClient("", "", nil)
	client.baseURL, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	service := FirewallZone{
		ID:        "id",
		Name:      "hello",
		IPAddress: IPList{"0.0.0.0/0"},
	}

	err = client.UpdateFirewallZone(service)

	if err != nil {
		t.Fatalf("firewall zone updating failed: %v", err)
	}

	if reqBody.Zone.ID != service.ID {
		t.Errorf("expected request to contain ID=%s; got %s", service.ID, reqBody.Zone.ID)
	}

	if reqBody.Zone.Name != service.Name {
		t.Errorf("expected request to contain Name=%s; got %s", service.Name, reqBody.Zone.Name)
	}

	if !reflect.DeepEqual(reqBody.Zone.IPAddress, service.IPAddress) {
		t.Errorf("expected request to contain IPAddress=%s; got %s", service.IPAddress, reqBody.Zone.IPAddress)
	}
}

func TestClient_DeleteFirewallZone(t *testing.T) {
	var err error

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "", http.StatusNoContent),
			t,
			http.MethodDelete,
			"/v1/firewall_zones/id",
			nil,
		),
	)
	defer ts.Close()

	client := NewClient("", "", nil)
	client.baseURL, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	err = client.DeleteFirewallZone("id")

	if err != nil {
		t.Fatalf("firewall zone deletion failed: %v", err)
	}
}
