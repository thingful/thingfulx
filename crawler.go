package thingfulx

// Crawler is the top level interface that implementing libraries must support
// to be a Thingful compatible parser. A Crawler is basically a simple
// container that returns a Fetcher and a Parser for a specific upstream
// provider identifed by the ID method on the Crawler instance.
type Crawler interface {
	// Fetcher returns an instantiated Fetcher instance for the upstream provider
	// this crawler knows how to retrieve data from.
	Fetcher() Fetcher

	// Parser returns an instantiated Parser instance for the upstream provider
	// this crawler knows how to parse the data from.
	Parser() Parser

	// Provider returns a unique identifier representing a single upstream source
	// of data. Typically this will be the Host component of the domain's URL,
	// but could be anything providing it is unique among crawlers.
	Provider() string
}
