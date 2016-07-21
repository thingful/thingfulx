package thingfulx

import (
	"net/http"
	"time"
)

const (
	// ClientToken is the key we use when setting a token on the
	// cleint clientInfo property
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

	// GetClientInfo is a utility method that allows the client to pass
	// extra information about its context to the fetcher that will
	// consume it.
	GetClientInfo() map[string]interface{}
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

// NewClientWithInfo is a constructor function that instantiates and returns a
// new client struct, setting the clientInfo property
func NewClientWithInfo(userAgent string, timeout time.Duration, clientInfo map[string]interface{}) Client {
	return &client{
		userAgent: userAgent,
		http: &http.Client{
			Timeout: timeout,
		},
		clientInfo: clientInfo,
	}
}

func (c *client) GetClientInfo() map[string]interface{} {
	if len(c.clientInfo) != 0 {
		return c.clientInfo
	}

	return nil
}

// client is our concrete implementation of the Client interface
type client struct {
	userAgent  string
	http       *http.Client
	clientInfo map[string]interface{}
}

// DoHTTP is a thin wrapper around the `Do` method on `net/http.Client`. It
// simply ensures we set the correct user agent header before making the
// request
func (c *client) DoHTTP(req *http.Request) (*http.Response, error) {
	// set the user agent here
	req.Header.Add("User-Agent", c.userAgent)

	return c.http.Do(req)
}
