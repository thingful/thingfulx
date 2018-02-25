package thingfulx

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/thingful/httpmock"
)

func TestDoHTTP(t *testing.T) {
	duration := time.Duration(1) * time.Second

	client := NewClient("thingful", duration)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterStubRequest(
		httpmock.NewStubRequest(
			"GET",
			"http://example.com/foo",
			httpmock.NewStringResponder(200, "ok"),
		),
	)

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

func TestGet(t *testing.T) {
	duration := time.Duration(1) * time.Second

	client := NewClient("thingful", duration)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterStubRequest(
		httpmock.NewStubRequest(
			"GET",
			"http://example.com",
			httpmock.NewStringResponder(200, "ok"),
		),
	)

	t.Run("with new Get method", func(t *testing.T) {
		got, err := client.Get("http://example.com")
		assert.Nil(t, err)

		assert.Equal(t, []byte("ok"), got)
	})

	t.Run("with old DoHTTPGetRequest method", func(t *testing.T) {
		got, err := client.DoHTTPGetRequest("http://example.com")
		assert.Nil(t, err)

		assert.Equal(t, []byte("ok"), got)
	})
}

func TestGetError(t *testing.T) {

	testcases := []struct {
		httpStatus int
		expected   error
	}{
		{
			httpStatus: http.StatusNotFound,
			expected:   ErrNotFound,
		},
		{
			httpStatus: http.StatusTeapot,
			expected:   NewErrUnexpectedResponse(strconv.Itoa(http.StatusTeapot)),
		},
	}

	for _, testcase := range testcases {
		duration := time.Duration(1) * time.Second

		client := NewClient("thingful", duration)

		httpmock.Reset()

		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterStubRequest(
			httpmock.NewStubRequest(
				"GET",
				"http://example.com",
				httpmock.NewBytesResponder(testcase.httpStatus, nil),
			),
		)

		_, err := client.Get("http://example.com")
		assert.Equal(t, testcase.expected, err)
	}
}
