package thingfulx

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
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

	defer func() {
		if cerr := resp.Body.Close(); err == nil && cerr != nil {
			err = cerr
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Unexpected error, got %#v", err)
	}

	assert.Equal(t, body, []byte("ok"))
}
