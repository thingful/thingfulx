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
	// make the request. We also pass in a TimeProvider which is used internally
	// by the fetcher to record the indexing time of the parser. This is to allow
	// for easier testing.
	Fetch(url, userAgent string, client *http.Client, timeProvider TimeProvider) ([]Thing, error)

	// Host is a function that should return a string containing the host this
	// fetcher is designed to handle. This should just be the host part and not
	// contain the protocol, or any path information, e.g. it should be
	// `api.thingspeak.com` rather than `https://api.thingspeak.com/channels`
	Host() string

	// URLS is a function that can be called by Pomelo, that then returns back a
	// slice of strings containing the minimum set of URLS that should be indexed
	// for this host.  This might be a single URL for hosts that publish
	// relatively few things, or it might be a whole bunch for hosts that have
	// pages of results, that Pomelo should iterate through. The key requirement
	// is that this function should return the smallest set of URLs required to
	// completely index the upstream data provider.
	//
	// We pass in a userAgent string to this function as it will typically have
	// to make some HTTP request to the upstream source, and this should be made
	// with the consistent user agent string. We also pass in an instantiated
	// http.Client object as this is designed to be shared.
	URLS(userAgent string, client *http.Client) ([]string, error)
}
