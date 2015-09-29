package thingfulx

import (
	"io"
	"io/ioutil"
	"net/http"
)

// Fetcher is the interface for things that know how to fetch some resource
// from somewhere handling any protocol weirdness or authentication
// requirements. The interface is agnostic in terms of protocol so an
// implementing class is free to perform whatever steps required to get the
// data.
type Fetcher interface {
	// Fetch is the method a Fetcher must implement which returns a byte array
	// retrieved from the given urlStr. It is passed in the request method, the
	// urlStr we want to fetch, a user agent string that clients must send with
	// outgoing requests, a handle to a resusable http Client instance, and an
	// optional body for POST requests.
	//
	// The semantics of method, urlStr and body are exactly the same as that for
	// the http.NewRequest constructor in the net/http package.
	//
	// We return an instantiated Response object which is a simple struct that
	// contains the HTTP status code and the body and possibly an error if the
	// fetch failed unexpectedly.  Note that an HTTP error from an upstream
	// source does not cause an error object to be returned.
	Fetch(method, urlStr, userAgent string, client *http.Client, body io.Reader) (*Response, error)
}

// NewFetcher creates an returns an instance of a simple fetcher. This fetcher
// is one that is able to fetch an arbitrary HTTP resource that doesn't require
// any authentication.
func NewFetcher() Fetcher {
	return &fetcher{}
}

// fetcher is an implementation of the fetcher interface for simple cases where
// no special authentication needs to be performed
type fetcher struct{}

// Fetch a resource from the upstream provider if we can.
func (f *fetcher) Fetch(method, urlStr, userAgent string, client *http.Client, body io.Reader) (*Response, error) {
	// create a new request object
	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		return nil, err
	}

	// We must set the user agent manually
	req.Header.Add("User-Agent", userAgent)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &Response{
		Body:       bytes,
		StatusCode: resp.StatusCode,
	}, nil
}
