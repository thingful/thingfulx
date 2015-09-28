package thingfulx

import (
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
	// retrieved from the given rawurl string. It is passed in the rawurl we want
	// to fetch, a user agent string we must send with the request, and a handle
	// to a resusable http Client instance
	Fetch(rawurl, userAgent string, client *http.Client) (*Response, error)
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
func (f *fetcher) Fetch(rawurl, userAgent string, client *http.Client) (*Response, error) {
	// create a new request object
	req, err := http.NewRequest("GET", rawurl, nil)
	if err != nil {
		return nil, err
	}

	// we must set the user agent manually currently
	req.Header.Add("User-Agent", userAgent)

	// make the requests
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
