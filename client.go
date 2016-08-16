package thingfulx

import (
	"net/http"
	"time"
)

const (
	// ClientToken is the key we use when setting a token on the
	// client clientInfo property
	ClientToken = "client-token"
)

// Client specifies an interface we will use for making outgoing requests
// within each Fetcher. Currently everything is HTTP, but that won't
// necessarily remain the case, so we add a layer of abstraction meaning we can
// add methods to our Client interface to support other protocols.
type Client interface {
	// DoHTTP is a thin wrapper around the `Do` method on `net/http.Client`. It
	// simply ensures we set the correct user agent header before making the
	// request
	DoHTTP(req *http.Request) (*http.Response, error)
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
	}
}

// client is our concrete implementation of the Client interface
type client struct {
	userAgent string
	http      *http.Client
}

// DoHTTP is a thin wrapper around the `Do` method on `net/http.Client`. It
// simply ensures we set the correct user agent header before making the
// request
func (c *client) DoHTTP(req *http.Request) (*http.Response, error) {
	// set the user agent here
	req.Header.Add("User-Agent", c.userAgent)

	return c.http.Do(req)
}
