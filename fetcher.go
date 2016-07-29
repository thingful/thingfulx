package thingfulx

import (
	"time"
)

// Fetcher is the main interface for things that know how to fetch some
// resource from somewhere handling any protocol weirdness or authentication
// requirements and return back a slice of parsed Thing instances. The
// interface is agnostic in terms of protocol so an implementing class is free
// to perform whatever steps required to get the data.
type Fetcher interface {
	// Fetch is the method that all crawlers must fulfil which actually perform
	// the task of going and getting data from some remote data provider, and
	// returning a slice of Thing objects extracted from that data source.
	// Certain providers might require multiple requests from different URLs to
	// actually retrieve the data so within this interface they are free to do
	// this.
	//
	// This function takes as parameters the url we want to get data from,
	// a Client object that the crawler will use to actually make the request. We
	// also pass in a TimeProvider which is used internally by the fetcher to
	// record the indexing time of the parser. This is to allow for easier
	// testing.
	Fetch(url string, client Client, context Context, timeProvider TimeProvider) ([]Thing, error)

	// Provider is a function returning an instantiated Provider object
	// describing the upstream data provider this particular fetcher can
	// indexing. This data structure contains the name of the provider, plus a
	// url.URL struct containing the base URL of the provider
	Provider() *Provider

	// URLS is a function that can be called by Pomelo, that then returns back a
	// slice of strings containing the minimum set of URLS that should be indexed
	// for this host.  This might be a single URL for hosts that publish
	// relatively few things, or it might be a whole bunch for hosts that have
	// pages of results, that Pomelo should iterate through. The key requirement
	// is that this function should return the smallest set of URLs required to
	// completely index the upstream data provider.
	//
	// We pass in as parameters a Client implementation which the Fetcher must
	// use to make any outgoing requests, plus a delay Duration. Requests should
	// not be made faster than the time interval specified by delay.
	URLS(client Client, delay time.Duration) ([]string, error)
}
