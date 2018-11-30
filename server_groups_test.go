package api

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestClient_ListServerGroups(t *testing.T) {
	var err error
	expectedResults := 2
	expectedID := "9981f162c2d611e680b17f1fb185c564"

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "server_groups_list", http.StatusOK),
			t,
			http.MethodGet,
			"/v1/groups",
			nil,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.BaseUrl, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	resp, err := client.ListServerGroups()

	if err != nil {
		t.Fatalf("server groups list failed: %v", err)
	}

	if resp.Count != expectedResults {
		t.Errorf("expected count to be %d; got %d", expectedResults, resp.Count)
	}

	if len(resp.Groups) != expectedResults {
		t.Errorf("expected %d groups; got %d", expectedResults, resp.Count)
	}

	if resp.Groups[0].ID != expectedID {
		t.Errorf("expected group 0 to have ID %s; got %s", expectedID, resp.Groups[0].ID)
	}
}

func TestClient_GetServerGroup(t *testing.T) {
	var err error
	expectedID := "0962bfa087bc01323e360670140ec224"

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "server_groups_get", http.StatusOK),
			t,
			http.MethodGet,
			"/v1/groups/id",
			nil,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.BaseUrl, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	resp, err := client.GetServerGroup("id")

	if err != nil {
		t.Fatalf("server groups get failed: %v", err)
	}

	if resp.Group.ID != expectedID {
		t.Errorf("expected group to have ID %s; got %s", expectedID, resp.Group.ID)
	}
}

func TestClient_CreateServerGroup(t *testing.T) {
	var err error
	reqBody := CreateServerGroupRequest{}

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "server_groups_get", http.StatusCreated),
			t,
			http.MethodPost,
			"/v1/groups",
			&reqBody,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.BaseUrl, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	name := "some-name"
	resp, err := client.CreateServerGroup(ServerGroup{Name: name})

	if err != nil {
		t.Fatalf("CSP account creating failed: %v", err)
	}

	if reqBody.Group.Name != name {
		t.Errorf("expected request to have group name %s; got %s", name, reqBody.Group.Name)
	}

	if len(resp.Group.ID) == 0 {
		t.Errorf("response did not include group ID")
	}
}

func TestClient_UpdateServerGroup(t *testing.T) {
	var err error
	reqBody := UpdateServerGroupRequest{}

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "server_groups_get", http.StatusCreated),
			t,
			http.MethodPut,
			"/v1/groups/some-id",
			&reqBody,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.BaseUrl, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	name := "another-name"
	err = client.UpdateServerGroup(ServerGroup{ID: "some-id", Name: name})

	if err != nil {
		t.Fatalf("CSP account creating failed: %v", err)
	}

	if reqBody.Group.Name != name {
		t.Errorf("expected request to have group name %s; got %s", name, reqBody.Group.Name)
	}
}

func TestClient_DeleteServerGroup(t *testing.T) {
	var err error

	ts := httptest.NewServer(
		requestValidatorTestHandler(
			jsonResponseTestHandler(t, "", http.StatusNoContent),
			t,
			http.MethodDelete,
			"/v1/groups/test",
			nil,
		),
	)
	defer ts.Close()

	client := NewClient("", "")
	client.BaseUrl, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	err = client.DeleteServerGroup("test")

	if err != nil {
		t.Fatalf("server group deletion failed: %v", err)
	}
}
