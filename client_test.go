package thingfulx

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/thingful/httpmock"
)

func newClient(permittedHosts ...string) Client {
	duration := time.Duration(1) * time.Second

	return NewClient("thingful", duration, permittedHosts...)
}

func TestDoHTTP(t *testing.T) {
	client := newClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterStubRequest(
		httpmock.NewStubRequest(
			"GET",
			"http://example.com/foo",
			httpmock.NewStringResponder(200, "ok"),
		).WithHeader(
			&http.Header{
				"User-Agent": []string{"thingful"},
			},
		),
	)

	httpmock.RegisterStubRequest(
		httpmock.NewStubRequest(
			"GET",
			"http://example.com/robots.txt",
			httpmock.NewStringResponder(200, "User-agent: *\nDisallow:"),
		).WithHeader(
			&http.Header{
				"User-Agent": []string{"thingful"},
			},
		),
	)

	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	assert.Nil(t, err)

	resp, err := client.DoHTTP(req)
	assert.Nil(t, err)

	if resp.Body != nil {
		defer func() {
			if cerr := resp.Body.Close(); err == nil && cerr != nil {
				err = cerr
			}
		}()
	}

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	assert.Equal(t, body, []byte("ok"))

	assert.Nil(t, httpmock.AllStubsCalled())
}

func TestDoHTTPWhenRobotstxtBlocks(t *testing.T) {
	client := newClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterStubRequest(
		httpmock.NewStubRequest(
			"GET",
			"http://example.com/robots.txt",
			httpmock.NewStringResponder(200, "User-agent: *\nDisallow: /"),
		).WithHeader(
			&http.Header{
				"User-Agent": []string{"thingful"},
			},
		),
	)

	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	assert.Nil(t, err)

	_, err = client.DoHTTP(req)
	if assert.Error(t, err, "Expected not permitted to fetch  error") {
		assert.Equal(t, err, ErrNotPermittedToFetch)
	}

	assert.Nil(t, httpmock.AllStubsCalled())
}

func TestDoHTTPWhenPermitted(t *testing.T) {
	client := newClient("example.com")

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterStubRequest(
		httpmock.NewStubRequest(
			"GET",
			"http://example.com/foo",
			httpmock.NewStringResponder(200, "ok"),
		).WithHeader(
			&http.Header{
				"User-Agent": []string{"thingful"},
			},
		),
	)

	httpmock.RegisterStubRequest(
		httpmock.NewStubRequest(
			"GET",
			"http://example.com/robots.txt",
			httpmock.NewStringResponder(200, "User-agent: *\nDisallow: /"),
		).WithHeader(
			&http.Header{
				"User-Agent": []string{"thingful"},
			},
		),
	)

	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	assert.Nil(t, err)

	resp, err := client.DoHTTP(req)
	assert.Nil(t, err)

	if resp.Body != nil {
		defer func() {
			if cerr := resp.Body.Close(); err == nil && cerr != nil {
				err = cerr
			}
		}()
	}

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	assert.Equal(t, body, []byte("ok"))

	assert.Nil(t, httpmock.AllStubsCalled())
}

func TestDoHTTPWhenRobotAgentCached(t *testing.T) {
	client := newClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterStubRequest(
		httpmock.NewStubRequest(
			"GET",
			"http://example.com/foo",
			httpmock.NewStringResponder(200, "ok"),
		).WithHeader(
			&http.Header{
				"User-Agent": []string{"thingful"},
			},
		),
	)

	httpmock.RegisterStubRequest(
		httpmock.NewStubRequest(
			"GET",
			"http://example.com/robots.txt",
			httpmock.NewStringResponder(200, "User-agent: *\nDisallow:"),
		).WithHeader(
			&http.Header{
				"User-Agent": []string{"thingful"},
			},
		),
	)

	httpmock.RegisterStubRequest(
		httpmock.NewStubRequest(
			"GET",
			"http://example.com/bar",
			httpmock.NewStringResponder(200, "ok"),
		).WithHeader(
			&http.Header{
				"User-Agent": []string{"thingful"},
			},
		),
	)

	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	assert.Nil(t, err)

	resp1, err := client.DoHTTP(req)
	assert.Nil(t, err)

	if resp1.Body != nil {
		defer func() {
			if cerr := resp1.Body.Close(); err == nil && cerr != nil {
				err = cerr
			}
		}()
	}

	body, err := ioutil.ReadAll(resp1.Body)
	assert.Nil(t, err)

	assert.Equal(t, body, []byte("ok"))

	req, err = http.NewRequest("GET", "http://example.com/bar", nil)
	assert.Nil(t, err)

	// TODO: figure out how to verify this - I can see from coverage report that
	// now we use the cached robot agent, but currently httpmock doesn't let you
	// assert how many times a mock is called.
	resp2, err := client.DoHTTP(req)
	assert.Nil(t, err)

	if resp2.Body != nil {
		defer func() {
			if cerr := resp2.Body.Close(); err == nil && cerr != nil {
				err = cerr
			}
		}()
	}

	body, err = ioutil.ReadAll(resp2.Body)
	assert.Nil(t, err)

	assert.Equal(t, body, []byte("ok"))

	assert.Nil(t, httpmock.AllStubsCalled())
}

func TestDoHTTPWhenRobotsNotFound(t *testing.T) {
	client := newClient()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterStubRequest(
		httpmock.NewStubRequest(
			"GET",
			"http://example.com/foo",
			httpmock.NewStringResponder(200, "ok"),
		).WithHeader(
			&http.Header{
				"User-Agent": []string{"thingful"},
			},
		),
	)

	httpmock.RegisterStubRequest(
		httpmock.NewStubRequest(
			"GET",
			"http://example.com/robots.txt",
			httpmock.NewStringResponder(404, ""),
		).WithHeader(
			&http.Header{
				"User-Agent": []string{"thingful"},
			},
		),
	)

	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	assert.Nil(t, err)

	resp, err := client.DoHTTP(req)
	assert.Nil(t, err)

	if resp.Body != nil {
		defer func() {
			if cerr := resp.Body.Close(); err == nil && cerr != nil {
				err = cerr
			}
		}()
	}

	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)

	assert.Equal(t, body, []byte("ok"))

	assert.Nil(t, httpmock.AllStubsCalled())
}
