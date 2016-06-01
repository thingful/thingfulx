package thingfulx

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
)

func TestDoHTTP(t *testing.T) {
	duration := time.Duration(1) * time.Second

	client := NewClient("thingful", duration)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://example.com/foo",
		httpmock.NewStringResponder(200, "ok"))

	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	if err != nil {
		t.Fatalf("Unexpected error, got %#v", err)
	}

	resp, err := client.DoHTTP(req)
	if err != nil {
		t.Fatalf("Unexpected error, got %#v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Unexpected error, got %#v", err)
	}

	if string(body) != "ok" {
		t.Errorf("Unexpected response body, expected 'ok', got '%s'", body)
	}

}
