package web

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"web-pet-project/internal/dbms/model"
)

var testRouter http.Handler
var testServer *httptest.Server

func init() {
	testRouter = setupRouter()
}

func TestContext_getIssuesHandler(t *testing.T) {

	testServer = httptest.NewServer(testRouter)
	defer testServer.Close()

	// Run server request
	endpoint := "/api/v1/issues"
	r, err := http.NewRequest(http.MethodGet, testServer.URL+endpoint, nil)
	if err != nil {
		t.Fatal(err)
	}

	r.Header.Add("content-type", "application/json")
	r.Header.Add("cache-control", "no-cache")

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		t.Fatal(err)
	}

	var issues []model.Issue
	err = json.NewDecoder(resp.Body).Decode(&issues)
	if err != nil {
		t.Fatal(err)
	}
	if len(issues) != 2 {
		t.Fatal("Bad issues count")
	}

	// Validate Response by status code
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected %d, got %d", http.StatusOK, resp.StatusCode)
	}

}
