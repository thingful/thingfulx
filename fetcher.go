// Package thingfulx provides a compatibility layer for implementing parsers
// for use with the Thingful indexing system. New upstream providers can be
// added by implementing the core Fetcher and Parser interfaces contained
// within this library. These libraries can then be imported and used by
// Thingful.
package thingfulx

import (
	"net/http"
)

// Fetcher is the main interface for things that know how to fetch some
// resource from somewhere handling any protocol weirdness or authentication
// requirements and return back a slice of parsed Thing instances. The
// interface is agnostic in terms of protocol so an implementing class is free
// to perform whatever steps required to get the data.
type Fetcher interface {
	// Fetch is the method that all crawler's must fulfil which actually perform
	// the task of going and getting data from some remote data provider, and
	// returning a slice of Thing objects extracted from that data source.
	// Certain providers might require multiple requests from different URLs to
	// actually retrieve the data so within this interface they are free to do
	// this.  This function takes as parameters the url we want to get data from,
	// a user agent string (passed in by the main crawler engine), and an
	// instantiated http.Client object that the crawler should use to actually
	// make the request.
	Fetch(url, userAgent string, client *http.Client) ([]Thing, error)
}
