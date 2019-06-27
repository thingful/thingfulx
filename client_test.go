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
