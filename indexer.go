package thingfulx

import (
	"context"
	"time"
)

// IndexerBuilder is type definition for the constructor functions that new
// Indexer must implement.
type IndexerBuilder func() (Indexer, error)

// URLData is a struct used to send back data from the URLS method. We now send
// back data asynchronously via a channel so we want to package our data we are
// actually interested in (i.e. the slice of strings), along with any error that
// might happen.
type URLData struct {
	// URLS is a slice of url strings representing a chunk of candidate resource
	// URLs for the target infrastructure.
	URLS []string

	// Err is used to allow the process to signal back an error to the caller. We
	// package it with the data so we can signal an error at any point.
	Err error
}

// Indexer is the main interface for things that know how to index and fetch
// resources from external data infrastructures and are thus responsible for
// handling any authentication requirements, parsing the returned data and
// generating normalized representations of IoT things to the caller. The
// interface attempts to be agnostic in terms of protocol so an implementing
// class is free to perform whatever steps required to get the data.
type Indexer interface {
	// UID returns a unique identifier for the data infrastructure indexed by
	// this Indexer. Good practice for this to be the same as the package name.
	UID() string

	// URLS returns the smallest set of URLs required to completely index the
	// upstream data provider. The returned values might be a single URL for hosts
	// that publish relatively few things that are all available via a single URL,
	// or it might be hundreds of thousands of values long for hosts that publish
	// data entirely on individual URLs.
	//
	// This function takes as parameters an instance of the Context interface for
	// request-scoped values or cancellation, a Client implementation which the
	// Indexer must use to make any outgoing requests, a delay Duration which
	// defines the minimum interval between consecutive requests to the
	// infrastructure. In addition the caller must pass in a channel by which the
	// function can return data to the caller. This is intended to allow the indexer
	// implementation to return chunks of data as it goes in order to support
	// infrastructures with hundreds of thousands of unique URLs to return.
	URLS(ctx context.Context, client Client, delay time.Duration, out chan<- URLData)

	// Fetch is our method that knows how to obtain raw data for an infrastructure,
	// and is therefore responsible for negotiating any authentication or protocol
	// weirdness when it comes to fetching data. It takes as parameters an instance
	// of the Context interface for request-scoped values and cancellation, the url
	// we want to get data from, a Client object that the indexer must use to
	// actually make the request, and a Clock to allow testing time related
	// functions. It returns the raw data from the infrastructure as a slice of
	// bytes for further parsing.
	Fetch(ctx context.Context, urlStr string, client Client, clock Clock) ([]byte, error)

	// Parse returns a slice of Channel objects extracted from that data source.
	// This function takes as parameters a slice of bytes representing the data
	// collected from the upstream data provider, the URL of the Channel being
	// indexed and a Clock instance used internally by the fetcher to record the
	// indexing time of the parser. This is to allow for easier testing. We
	// separate Parsing from Fetching as we have some systems that provide data to
	// us without having to be fetched from a remote HTTP source.
	Parse(rawData []byte, urlStr string, clock Clock) ([]*Channel, error)
}
