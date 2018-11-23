package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestClient_ListServers(t *testing.T) {
	var err error
	expectedResults := 1
	expectedID := "d54be8ca88ea11e8800f753bfb4b1x97"

	ts := httptest.NewServer(jsonResponseTestHandler("servers_list", t, true))
	defer ts.Close()

	client := NewClient("", "")
	client.BaseUrl, err = url.Parse(ts.URL)

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

	ts := httptest.NewServer(jsonResponseTestHandler("server_get", t, true))
	defer ts.Close()

	client := NewClient("", "")
	client.BaseUrl, err = url.Parse(ts.URL)

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

	ts := httptest.NewServer(authTestHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertRequest(r, t, http.MethodDelete, "/v1/servers/test")
		w.WriteHeader(http.StatusNoContent)
	}), t))
	defer ts.Close()

	client := NewClient("", "")
	client.BaseUrl, err = url.Parse(ts.URL)

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

	ts := httptest.NewServer(authTestHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertRequest(r, t, http.MethodPut, "/v1/servers/test")

		b, err := ioutil.ReadAll(r.Body)

		if err != nil {
			t.Fatalf("reading body failed: %v", err)
		}

		reqData := RetireServerRequest{}
		if err := json.Unmarshal(b, &reqData); err != nil {
			t.Fatalf("unmarshalling body failed: %v", err)
		}

		if !reqData.Server.Retire {
			t.Error("retirement should be set to true")
		}

		w.WriteHeader(http.StatusNoContent)
	}), t))
	defer ts.Close()

	client := NewClient("", "")
	client.BaseUrl, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	err = client.RetireServer("test")

	if err != nil {
		t.Fatalf("server retirement failed: %v", err)
	}
}

func TestClient_MoveServer(t *testing.T) {
	var err error

	testGroupID := "group_id"
	ts := httptest.NewServer(authTestHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assertRequest(r, t, http.MethodPut, "/v1/servers/test")

		b, err := ioutil.ReadAll(r.Body)

		if err != nil {
			t.Fatalf("reading body failed: %v", err)
		}

		reqData := MoveServerRequest{}
		if err := json.Unmarshal(b, &reqData); err != nil {
			t.Fatalf("unmarshalling body failed: %v", err)
		}

		if reqData.Server.GroupID != testGroupID {
			t.Errorf("expected group id %s; got %s", testGroupID, reqData.Server.GroupID)
		}

		w.WriteHeader(http.StatusNoContent)
	}), t))
	defer ts.Close()

	client := NewClient("", "")
	client.BaseUrl, err = url.Parse(ts.URL)

	if err != nil {
		t.Fatalf("cannot parse url %s: %v", ts.URL, err)
	}

	err = client.MoveServer("test", testGroupID)

	if err != nil {
		t.Fatalf("server movement failed: %v", err)
	}
}
