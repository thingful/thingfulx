package thingfulx

import (
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

	_, err = client.DoHTTP(req)
	if err != nil {
		t.Fatalf("Unexpected error, got %#v", err)
	}
}
