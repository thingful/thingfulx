package thingfulx

import (
	"net/http"
)

// Fetcher is the interface for things that know how to fetch some resource
// from somewhere handling any protocol weirdness or authentication
// requirements. The interface is agnostic in terms of protocol so an
// implementing class is free to perform whatever steps required to get the
// data.
type Fetcher interface {
	// Fetch is the method a Fetcher must implement which returns a byte array
	// retrieved from the given rawurl string. It is passed in the rawurl we want
	// to fetch, a user agent string we must send with the request, and a handle
	// to a resusable http Client instance
	Fetch(rawurl, userAgent string, client *http.Client) (*Response, error)
}
