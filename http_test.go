package thingfulx

import (
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/thingful/httpmock"
)

func TestHTTPGetRequest(t *testing.T) {
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

	got, err := HTTPGetRequest("http://example.com", client)
	assert.Nil(t, err)

	assert.Equal(t, []byte("ok"), got)
}

func TestHTTPGetRequestError(t *testing.T) {

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

		_, err := HTTPGetRequest("http://example.com", client)
		assert.Equal(t, testcase.expected, err)
	}
}
