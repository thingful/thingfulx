// Package thingfulx provides a compatibility layer for implementing parsers
// for use with the Thingful indexing system. New upstream providers can be
// added by implementing the core Fetcher and Parser interfaces contained
// within this library. These libraries can then be imported and used by
// Thingful.
package thingfulx

// Crawler is a struct containing a Fetcher and a Parser that together know how
// to fetch data from some upstream provider and parse it.
type Crawler struct {
	// Fetcher is an embedded Fetcher instance
	Fetcher

	// Parser is an embedded Parser instance
	Parser

	// Provider returns a unique string identifier representing the upstream data
	// provider this Crawler knows about.
	Provider string
}
