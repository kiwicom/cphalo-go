package api

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestClient_ListServers(t *testing.T) {
	var err error
	expectedResults := 1
	expectedID := "d54be8ca88ea11e8800f753bfb4b1x97"

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "servers_list", http.StatusOK),
			t,
			http.MethodGet,
			"/v1/servers",
			nil,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.baseURL, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	resp, err := client.ListServers()

	if err != nil {
		t.Fatalf("servers list failed: %v", err)
	}

	if resp.Count != expectedResults {
		t.Errorf("expected count to be %d; got %d", expectedResults, resp.Count)
	}

	if len(resp.Servers) != expectedResults {
		t.Errorf("expected %d servers; got %d", expectedResults, resp.Count)
	}

	if resp.Servers[0].ID != expectedID {
		t.Errorf("expected server 0 to have ID %s; got %s", expectedID, resp.Servers[0].ID)
	}
}

func TestClient_GetServer(t *testing.T) {
	var err error
	expectedID := "3958fe0c08e511e7819335b35e8ba368"

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "server_get", http.StatusOK),
			t,
			http.MethodGet,
			"/v1/servers/id",
			nil,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.baseURL, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	resp, err := client.GetServer("id")

	if err != nil {
		t.Fatalf("server groups get failed: %v", err)
	}

	if resp.Server.ID != expectedID {
		t.Errorf("expected server to have ID %s; got %s", expectedID, resp.Server.ID)
	}

	expectedSGID := "b864e2204f72012f94c9404038a8a7aa"
	if resp.Server.GroupID != expectedSGID {
		t.Errorf("expected server to be in ServerGroupID %s; got %s", expectedID, expectedSGID)
	}
}

func TestClient_DeleteServer(t *testing.T) {
	var err error

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "", http.StatusNoContent),
			t,
			http.MethodDelete,
			"/v1/servers/test",
			nil,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.baseURL, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	err = client.DeleteServer("test")

	if err != nil {
		t.Fatalf("server deletion failed: %v", err)
	}
}

func TestClient_RetireServer(t *testing.T) {
	var err error

	body := RetireServerRequest{}

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "", http.StatusNoContent),
			t,
			http.MethodPut,
			"/v1/servers/test",
			&body,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.baseURL, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	err = client.RetireServer("test")

	if err != nil {
		t.Fatalf("server retirement failed: %v", err)
	}

	if !body.Server.Retire {
		t.Error("retirement should be set to true")
	}
}

func TestClient_MoveServer(t *testing.T) {
	var err error

	testGroupID := "group_id"

	body := MoveServerRequest{}

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "", http.StatusNoContent),
			t,
			http.MethodPut,
			"/v1/servers/test",
			&body,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.baseURL, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	err = client.MoveServer("test", testGroupID)

	if err != nil {
		t.Fatalf("server movement failed: %v", err)
	}

	if body.Server.GroupID != testGroupID {
		t.Errorf("expected group id %s; got %s", testGroupID, body.Server.GroupID)
	}
}
