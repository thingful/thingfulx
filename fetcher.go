package thingfulx

import (
	"net/http"
)

// Fetcher is the interface for things that know how to fetch some resource
// from somewhere handling any authentication requirements
type Fetcher interface {
	// Fetch is the method a Fetcher must implement which returns a byte array
	// retrieved from the given rawurl string. It is passed in the rawurl we want
	// to fetch, a user agent string we must send with the request, and a handle
	// to a resusable http Client instance
	Fetch(rawurl, userAgent string, client *http.Client) (*Response, error)

	// Host returns the host that a Fetcher knows how to handle. This must be
	// just the Host part of the URL, i.e. what will be returned from the Host
	// accessor on a URL instance.
	Host() string
}
