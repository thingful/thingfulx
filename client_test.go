package thingfulx

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/thingful/simular"
)

//func TestDoHTTP(t *testing.T) {
//	duration := time.Duration(1) * time.Second
//
//	client := NewClient("thingful", duration)
//
//	httpmock.Activate()
//	defer httpmock.DeactivateAndReset()
//
//	httpmock.RegisterStubRequest(
//		httpmock.NewStubRequest(
//			"GET",
//			"http://example.com/foo",
//			httpmock.NewStringResponder(200, "ok"),
//		),
//	)
//
//	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
//	if err != nil {
//		t.Fatalf("Unexpected error, got %#v", err)
//	}
//
//	resp, err := client.DoHTTP(req)
//	if err != nil {
//		t.Fatalf("Unexpected error, got %#v", err)
//	}
//
//	defer func() {
//		if cerr := resp.Body.Close(); err == nil && cerr != nil {
//			err = cerr
//		}
//	}()
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		t.Fatalf("Unexpected error, got %#v", err)
//	}
//
//	assert.Equal(t, body, []byte("ok"))
//}

func TestGet(t *testing.T) {
	timeout := time.Duration(1) * time.Second

	client := NewClient("thingful", timeout)

	simular.Activate()
	defer simular.DeactivateAndReset()

	simular.RegisterStubRequests(
		simular.NewStubRequest(
			http.MethodGet,
			"http://example.com",
			simular.NewStringResponder(200, "ok"),
		),
	)

	ctx := context.Background()

	t.Run("with Get method", func(t *testing.T) {
		got, err := client.GetHTTP(ctx, "http://example.com")
		assert.Nil(t, err)

		assert.Equal(t, []byte("ok"), got)
	})

	assert.Nil(t, simular.AllStubsCalled())
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

	ctx := context.Background()

	for _, testcase := range testcases {
		duration := time.Duration(1) * time.Second

		client := NewClient("thingful", duration)

		simular.Reset()

		simular.Activate()
		defer simular.DeactivateAndReset()

		simular.RegisterStubRequests(
			simular.NewStubRequest(
				http.MethodGet,
				"http://example.com",
				simular.NewBytesResponder(testcase.httpStatus, nil),
			),
		)

		_, err := client.GetHTTP(ctx, "http://example.com")
		assert.Equal(t, testcase.expected, err)
	}
}

func TestPost(t *testing.T) {
	timeout := time.Duration(1) * time.Second

	client := NewClient("thingful", timeout)

	simular.Activate()
	defer simular.DeactivateAndReset()

	simular.RegisterStubRequests(
		simular.NewStubRequest(
			http.MethodPost,
			"http://example.com",
			simular.NewStringResponder(200, "ok"),
			simular.WithBody(
				strings.NewReader(`{"msg":"query"}`),
			),
			simular.WithHeader(
				&http.Header{
					"Content-Type": []string{"application/json"},
					"User-Agent":   []string{"thingful"},
				},
			),
		),
	)

	ctx := context.Background()

	got, err := client.PostHTTP(ctx, "http://example.com", "application/json", strings.NewReader(`{"msg":"query"}`))
	assert.Nil(t, err)

	assert.Equal(t, []byte("ok"), got)
}
