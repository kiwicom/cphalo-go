package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestClient_ListServerGroups(t *testing.T) {
	var err error
	expectedResults := 2
	expectedID := "9981f162c2d611e680b17f1fb185c564"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadFile("example_responses/server_groups_list.json")

		if err != nil {
			t.Fatalf("cannot read file: %v", err)
		}

		fmt.Fprint(w, string(b))
	}))
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

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadFile("example_responses/server_groups_get.json")

		if err != nil {
			t.Fatalf("cannot read file: %v", err)
		}

		fmt.Fprint(w, string(b))
	}))
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