package api

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestClient_ListFirewallServices(t *testing.T) {
	var err error
	expectedResults := 27
	expectedID := "ea3fe1309956012ee2989ea98as989as"

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "firewall_services_list", http.StatusOK),
			t,
			http.MethodGet,
			"/v1/firewall_services",
			nil,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.BaseUrl, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	resp, err := client.ListFirewallServices()

	if err != nil {
		t.Fatalf("firewall services list failed: %v", err)
	}

	if resp.Count != expectedResults {
		t.Errorf("expected count to be %d; got %d", expectedResults, resp.Count)
	}

	if len(resp.Services) != expectedResults {
		t.Errorf("expected %d firewall services; got %d", expectedResults, resp.Count)
	}

	if resp.Services[0].ID != expectedID {
		t.Errorf("expected firewall service 0 to have ID %s; got %s", expectedID, resp.Services[0].ID)
	}
}

func TestClient_GetFirewallService(t *testing.T) {
	var err error
	expectedID := "ea3fe1309956012ee2989ea98as989as"

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "firewall_services_get", http.StatusOK),
			t,
			http.MethodGet,
			"/v1/firewall_services/id",
			nil,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.BaseUrl, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	resp, err := client.GetFirewallService("id")

	if err != nil {
		t.Fatalf("firewall services get failed: %v", err)
	}

	if resp.Service.ID != expectedID {
		t.Errorf("expected firewall service to have ID %s; got %s", expectedID, resp.Service.ID)
	}
}

func TestClient_CreateFirewallService(t *testing.T) {
	var err error
	reqBody := CreateFirewallServiceRequest{}

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "firewall_services_get", http.StatusOK),
			t,
			http.MethodPost,
			"/v1/firewall_services",
			&reqBody,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.BaseUrl, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	service := FirewallService{
		ID:       "id",
		Name:     "hello",
		Protocol: "UDP",
		Port:     "53",
	}

	resp, err := client.CreateFirewallService(service)

	if err != nil {
		t.Fatalf("Firewall service creating failed: %v", err)
	}

	expectedID := "ea3fe1309956012ee2989ea98as989as"
	if resp.Service.ID != expectedID {
		t.Errorf("expected response to containt ID=%s; got %s", expectedID, resp.Service.ID)
	}

	if reqBody.Service.ID != service.ID {
		t.Errorf("expected request to contain ID=%s; got %s", service.ID, reqBody.Service.ID)
	}

	if reqBody.Service.Name != service.Name {
		t.Errorf("expected request to contain Name=%s; got %s", service.Name, reqBody.Service.Name)
	}

	if reqBody.Service.Protocol != service.Protocol {
		t.Errorf("expected request to contain Protocol=%s; got %s", service.Protocol, reqBody.Service.Protocol)
	}

	if reqBody.Service.Port != service.Port {
		t.Errorf("expected request to contain Port=%s; got %s", service.Port, reqBody.Service.Port)
	}
}

func TestClient_UpdateFirewallService(t *testing.T) {
	var err error
	reqBody := UpdateFirewallServiceRequest{}

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "", http.StatusOK),
			t,
			http.MethodPut,
			"/v1/firewall_services/id",
			&reqBody,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.BaseUrl, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	service := FirewallService{
		ID:       "id",
		Name:     "hello",
		Protocol: "TCP",
		Port:     "6667",
	}

	err = client.UpdateFirewallService(service)

	if err != nil {
		t.Fatalf("Firewall service updating failed: %v", err)
	}

	if reqBody.Service.ID != service.ID {
		t.Errorf("expected request to contain ID=%s; got %s", service.ID, reqBody.Service.ID)
	}

	if reqBody.Service.Name != service.Name {
		t.Errorf("expected request to contain Name=%s; got %s", service.Name, reqBody.Service.Name)
	}

	if reqBody.Service.Protocol != service.Protocol {
		t.Errorf("expected request to contain Protocol=%s; got %s", service.Protocol, reqBody.Service.Protocol)
	}

	if reqBody.Service.Port != service.Port {
		t.Errorf("expected request to contain Port=%s; got %s", service.Port, reqBody.Service.Port)
	}
}

func TestClient_DeleteFirewallService(t *testing.T) {
	var err error

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "", http.StatusNoContent),
			t,
			http.MethodDelete,
			"/v1/firewall_services/id",
			nil,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.BaseUrl, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	err = client.DeleteFirewallService("id")

	if err != nil {
		t.Fatalf("Firewall service deletion failed: %v", err)
	}
}
