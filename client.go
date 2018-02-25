package thingfulx

import (
	"io/ioutil"
	"net/http"
	"time"
)

// ContextKey is a type alias used for context keys to avoid any issues wth clashing keys
type ContextKey string

const (
	// ClientToken is the key we use when setting a token on a context when
	// making a request
	ClientToken = ContextKey("client-token")
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

	// Get is a helper function on the client interface for the most common usage
	// of the Client, making a simple Get request for a given URL, and returning
	// a slice of bytes containing the HTTP response. This uses client.DoHTTP
	// internally so that it ensures the correct user agent string is sent.
	Get(urlStr string) ([]byte, error)

	// DoHTTPGetRequest is a helper function on the client interface for the most
	// common usage of the Client, making a simple Get request for a given URL,
	// and returning a slice of bytes containing the HTTP response. This uses
	// client.DoHTTP internally so that it ensures the correct user agent string
	// is sent.
	//
	// Deprecated: Use Client.Get instead
	DoHTTPGetRequest(urlString string) ([]byte, error)
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

// Get is a helper function on the client interface for the most common usage
// of the Client, making a simple Get request for a given URL, and returning a
// slice of bytes containing the HTTP response. This uses client.DoHTTP
// internally to ensure that the correct user agent string is sent.
func (c *client) Get(urlStr string) (b []byte, err error) {
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.DoHTTP(req)
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

// DoHTTPGetRequest is a helper function on the client interface for the most
// common usage of the Client, making a simple Get request for a given URL, and
// returning a slice of bytes containing the HTTP response. This uses
// client.DoHTTP internally to ensure that the correct user agent string is
// sent.
//
// Deprecated: Use Client.Get instead
func (c *client) DoHTTPGetRequest(urlStr string) ([]byte, error) {
	return c.Get(urlStr)
}
