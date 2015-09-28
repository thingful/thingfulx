package thingfulx

import (
	"github.com/jarcoal/httpmock"
	"net/http"
	"testing"
)

func TestNewFetcher(t *testing.T) {
}

func TestFetchValid(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	fetcher := NewFetcher()

	client := &http.Client{}

	testcases := []struct {
		method         string
		url            string
		respStatusCode int
		respBody       string
		response       *Response
	}{
		{
			method:         "GET",
			url:            "http://example.com/thing",
			respStatusCode: 200,
			respBody:       "valid-data",
			response: &Response{
				StatusCode: 200,
				Body:       []byte("valid-data"),
			},
		},
	}

	for _, testcase := range testcases {
		httpmock.Reset()
		httpmock.RegisterResponder(
			testcase.method,
			testcase.url,
			httpmock.NewStringResponder(testcase.respStatusCode, testcase.respBody),
		)

		response, err := fetcher.Fetch(testcase.url, "Thingbot", client)
		if err != nil {
			t.Errorf("Unexpected error when calling Fetch, got '%s'", err.Error())
		}

		if response.StatusCode != testcase.respStatusCode {
			t.Errorf("Unexpected status code, expected '%d', got '%d'", testcase.respStatusCode, response.StatusCode)
		}

		if string(response.Body) != testcase.respBody {
			t.Errorf("Unexpected response body, expected '%s', got '%s'", testcase.respBody, string(response.Body))
		}
	}
}

func TestFetchError(t *testing.T) {
	fetcher := NewFetcher()

	client := &http.Client{}

	testcases := []struct {
		method string
		url    string
	}{
		{
			method: "GET",
			url:    "",
		},
		{
			method: "GET",
			url:    "%",
		},
	}

	for _, testcase := range testcases {
		_, err := fetcher.Fetch(testcase.url, "Thingbot", client)
		if err == nil {
			t.Errorf("Expected error for: '%s %s'", testcase.method, testcase.url)
		}
	}
}
