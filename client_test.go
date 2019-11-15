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

func TestGetNoRobots(t *testing.T) {
	timeout := time.Duration(500) * time.Millisecond

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

	got, err := client.GetHTTP(ctx, "http://example.com")
	assert.Nil(t, err)

	assert.Equal(t, []byte("ok"), got)

	assert.Nil(t, simular.AllStubsCalled())
}

func TestGetPermissiveRobots(t *testing.T) {
	timeout := time.Duration(500) * time.Millisecond

	// manually construct instance (not interface), so we can inspect state of
	// robots after use
	cl := &client{
		userAgent: "thingful",
		http: &http.Client{
			Timeout: timeout,
		},
		robots: newRobots(),
	}

	simular.Activate()
	defer simular.DeactivateAndReset()

	simular.RegisterStubRequests(
		simular.NewStubRequest(
			http.MethodGet,
			"http://example.com/robots.txt",
			simular.NewStringResponder(200, "User-Agent: *\nDisallow:"),
		),
		simular.NewStubRequest(
			http.MethodGet,
			"http://example.com/some/resource/1",
			simular.NewStringResponder(200, "ok"),
		),
		simular.NewStubRequest(
			http.MethodGet,
			"http://example.com/some/resource/2",
			simular.NewStringResponder(200, "ok"),
		),
	)

	ctx := context.Background()

	got, err := cl.GetHTTP(ctx, "http://example.com/some/resource/1")
	assert.Nil(t, err)
	assert.Equal(t, []byte("ok"), got)

	assert.Len(t, cl.robots.cache, 1)

	got, err = cl.GetHTTP(ctx, "http://example.com/some/resource/2")
	assert.Nil(t, err)
	assert.Equal(t, []byte("ok"), got)

	assert.Len(t, cl.robots.cache, 1)

	assert.Nil(t, simular.AllStubsCalled())
}

func TestGetWithRestrictiveRobotsTxt(t *testing.T) {
	timeout := time.Duration(500) * time.Millisecond

	client := NewClient("thingful", timeout)

	simular.Activate()
	defer simular.DeactivateAndReset()

	simular.RegisterStubRequests(
		simular.NewStubRequest(
			http.MethodGet,
			"http://example.com/robots.txt",
			simular.NewStringResponder(200, "User-agent: friendly\nDisallow:\n\nUser-agent: thingful\nDisallow: /some/"),
		),
	)

	ctx := context.Background()

	_, err := client.GetHTTP(ctx, "http://example.com/some/resource")
	assert.NotNil(t, err)
	assert.Equal(t, ErrRobotsForbidden, err)

	assert.Nil(t, simular.AllStubsCalled())
}

func TestGetInvalidRobots(t *testing.T) {
	timeout := time.Duration(500) * time.Millisecond

	client := NewClient("thingful", timeout)

	simular.Activate()
	defer simular.DeactivateAndReset()

	simular.RegisterStubRequests(
		simular.NewStubRequest(
			http.MethodGet,
			"http://example.com/robots.txt",
			simular.NewStringResponder(603, "User-Agent: *\nDisal low ALLOWsd"),
		),
		simular.NewStubRequest(
			http.MethodGet,
			"http://example.com",
			simular.NewStringResponder(200, "ok"),
		),
	)

	ctx := context.Background()

	got, err := client.GetHTTP(ctx, "http://example.com")
	assert.Nil(t, err)

	assert.Equal(t, []byte("ok"), got)

	assert.Nil(t, simular.AllStubsCalled())
}

func TestGetMissingGroup(t *testing.T) {
	timeout := time.Duration(500) * time.Millisecond

	client := NewClient("thingful", timeout)

	simular.Activate()
	defer simular.DeactivateAndReset()

	simular.RegisterStubRequests(
		simular.NewStubRequest(
			http.MethodGet,
			"http://example.com/robots.txt",
			simular.NewStringResponder(200, "User-Agent: googlebot\nDisallow:"),
		),
		simular.NewStubRequest(
			http.MethodGet,
			"http://example.com",
			simular.NewStringResponder(200, "ok"),
		),
	)

	ctx := context.Background()

	got, err := client.GetHTTP(ctx, "http://example.com")
	assert.Nil(t, err)

	assert.Equal(t, []byte("ok"), got)

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

		assert.Nil(t, simular.AllStubsCalled())
	}
}

func TestPostNoRobots(t *testing.T) {
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

	assert.Nil(t, simular.AllStubsCalled())
}

func TestPostPermissiveRobots(t *testing.T) {
	timeout := time.Duration(1) * time.Second

	client := NewClient("thingful", timeout)

	simular.Activate()
	defer simular.DeactivateAndReset()

	simular.RegisterStubRequests(
		simular.NewStubRequest(
			http.MethodGet,
			"http://example.com/robots.txt",
			simular.NewStringResponder(200, "User-Agent: *\nDisallow:"),
		),
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

	assert.Nil(t, simular.AllStubsCalled())
}

func TestPostError(t *testing.T) {
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
				http.MethodPost,
				"http://example.com",
				simular.NewStringResponder(testcase.httpStatus, "not ok"),
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

		_, err := client.PostHTTP(ctx, "http://example.com", "application/json", strings.NewReader(`{"msg":"query"}`))
		assert.Equal(t, testcase.expected, err)

		assert.Nil(t, simular.AllStubsCalled())
	}
}
