package thingfulx

import (
	"github.com/jarcoal/httpmock"
	"io"
	"net/http"
	"strings"
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
		urlStr         string
		body           io.Reader
		respStatusCode int
		respBody       string
		response       *Response
	}{
		{
			method:         "GET",
			urlStr:         "http://example.com/thing",
			body:           nil,
			respStatusCode: 200,
			respBody:       "valid-data",
			response: &Response{
				StatusCode: 200,
				Body:       []byte("valid-data"),
			},
		},
		{
			method:         "GET",
			urlStr:         "http://example.com/thing",
			body:           nil,
			respStatusCode: 404,
			respBody:       "not-found",
			response: &Response{
				StatusCode: 404,
				Body:       []byte("not-found"),
			},
		},
		{
			method:         "POST",
			urlStr:         "http://example.com/thing",
			body:           strings.NewReader("foo=bar"),
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
			testcase.urlStr,
			httpmock.NewStringResponder(testcase.respStatusCode, testcase.respBody),
		)

		response, err := fetcher.Fetch(testcase.method, testcase.urlStr, "Thingbot", client, testcase.body)
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
		urlStr string
		body   io.Reader
	}{
		{
			method: "GET",
			urlStr: "",
			body:   nil,
		},
		{
			method: "GET",
			urlStr: "%",
			body:   nil,
		},
	}

	for _, testcase := range testcases {
		_, err := fetcher.Fetch(testcase.method, testcase.urlStr, "Thingbot", client, testcase.body)
		if err == nil {
			t.Errorf("Expected error for: '%s %s'", testcase.method, testcase.urlStr)
		}
	}
}
