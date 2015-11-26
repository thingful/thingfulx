package thingfulx

import (
	"time"
)

// Thing is the top level datastructure representing an individual thing as
// indexed on some upstream data provider.
type Thing struct {
	// Title is a short descriptive lable for the thing
	Title string

	// Description is a long form description about the thing
	Description string

	// Category is a thingfulx.category instance - choose from the declared categories
	Category Category

	// Webpage is a URL to a human intelligible page about the thing
	Webpage string

	// URI is a unique identifier of the Thing. This does not have to be the URL
	// from which data for this thing may be actually fetched as some things only
	// publish their data as a single file containing multiple resources. In fact
	// this doesn't even have to be a URL - it could be a URN style identifier.
	// However by convention we've tended to make it into a URL. If the thing has
	// no URL of it's own we've tended to synthesize one by for example adding a
	// unique HTTP anchor to the block URL, e.g. http://example.com#abc123
	URI string

	// IndexedAt is a Time object indicating the timestamp when the thing was
	// indexed by Thingful
	IndexedAt time.Time

	// Lng is the longitude of the Thing, represented as a decimal value
	Lng float64

	// Lat is the latitude of the Thing, represented as a decimal value
	Lat float64

	// Provider is a data structure containing information about the upstream
	// data provider.
	Provider Provider

	// Visibility indicates what level of access the upstream data provider
	// offers, Closed, Shared or Open.
	Visibility Visibility

	// DataURL is a string containing the actual URL we go to fetch data from for
	// this thing
	DataURL string

	// Channels are a slice of Channel objects which capture the actual data of
	// the thing
	Channels []Channel
}
