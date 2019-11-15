package thingfulx

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// ContextKey is a type alias used for context keys to avoid any issues wth
// clashing keys
type ContextKey string

const (
	// ClientToken is the key we use when setting a token on a context when making
	// a request. Note it is the responsibility of the implementing indexer to
	// determine how to send this token value to the data infrastructure (i.e.
	// Authorization header, API Key in the query string or a header).
	ClientToken = ContextKey("client-token")
)

// Client specifies an interface we will use for making outgoing requests
// within each Fetcher. Currently everything is HTTP, but that won't
// necessarily remain the case, so we add a layer of abstraction meaning we can
// add methods to our Client interface to support other protocols.
type Client interface {
	// DoHTTP is a thin wrapper around the `Do` method on `net/http.Client`. It
	// simply ensures we set the correct user agent header before making the
	// request. Note that unlike the wrapped standard library method we pass
	// context explicitly as the first parameter.
	DoHTTP(ctx context.Context, req *http.Request) (*http.Response, error)

	// Get is a helper function on the client interface for the most common usage
	// of the Client, making a simple Get request for a given URL, and returning a
	// slice of bytes containing the HTTP response. This uses client.DoHTTP
	// internally so that it ensures the correct user agent string is sent. Note
	// that unlike the wrapped standard library method we pass context explicitly
	// as the first parameter.
	GetHTTP(ctx context.Context, urlStr string) ([]byte, error)

	// PostHTTP defines a function for sending a POST request to a data
	// infrastructure in order to obtain data. Typically we send GET requests, but
	// some infrastructures may require POST requests. This function returns a
	// slice of bytes for any HTTP OK response, and an error for anything else. If
	// you require finer control of this you should use the DoHTTP method directly
	// which returns the http.Response for processing as required. Note that unlike
	// the wrapped standard library method we pass context explicitly as the first
	// parameter.
	PostHTTP(ctx context.Context, urlStr, contentType string, body io.Reader) ([]byte, error)
}

// NewClient is a constructor function that instantiates and returns a new
// client struct, ensuring it sets `timeout` properly on the returned
// http.Client
func NewClient(userAgent string, timeout time.Duration) Client {
	return &client{
		userAgent: userAgent,
		http: &http.Client{
			Timeout: timeout,
		},
		robots: newRobots(),
	}
}

// client is our concrete implementation of the Client interface
type client struct {
	userAgent string
	http      *http.Client
	robots    *robots
}

// DoHTTP is a thin wrapper around the `Do` method on `net/http.Client`. It
// simply ensures we set the correct user agent header before making the
// request. We return an http.Response struct or an error.
func (c *client) DoHTTP(ctx context.Context, req *http.Request) (*http.Response, error) {
	if !c.robots.permitted(ctx, req.URL, c.userAgent, c.http) {
		return nil, ErrRobotsForbidden
	}

	// set the user agent here
	req.Header.Add("User-Agent", c.userAgent)

	// add our context to the request - this allows us to pass down any
	// cancellation to where it needs to go.
	req = req.WithContext(ctx)

	return c.http.Do(req)
}

// GetHTTP is a helper function on the client interface for the most common
// usage of the Client, making a simple Get request for a given URL, and
// returning a slice of bytes containing the HTTP response. This uses
// client.DoHTTP internally to ensure that the correct user agent string is
// sent. Note unlike the default http client interface we expose context as a
// parameter in recommended style for context usage. We return a slice of bytes
// or an error.
func (c *client) GetHTTP(ctx context.Context, urlStr string) (b []byte, err error) {
	req, err := http.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.DoHTTP(ctx, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if cerr := resp.Body.Close(); err == nil && cerr != nil {
			err = cerr
		}
	}()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, ErrNotFound
		}
		return nil, NewErrUnexpectedResponse(resp.Status)
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, err
}

// PostHTTP is our implementation of the interface method for fetching data via
// HTTP POST requests. It takes as parameters a context, the url to which data
// should be posted and its content type, as well as the actual request body
// wrapped in an io.Reader. Note that unlike the wrapped standard library method
// we pass context explicitly as the first parameter. We return a slice of bytes
// or an error.
func (c *client) PostHTTP(ctx context.Context, urlStr, contentType string, body io.Reader) (b []byte, err error) {
	req, err := http.NewRequest(http.MethodPost, urlStr, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	resp, err := c.DoHTTP(ctx, req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if cerr := resp.Body.Close(); err == nil && cerr != nil {
			err = cerr
		}
	}()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, ErrNotFound
		}
		return nil, NewErrUnexpectedResponse(resp.Status)
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, err
}
