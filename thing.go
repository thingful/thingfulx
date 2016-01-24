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

	// Metadata is a comma separated tags - additional metadata about the Thing,
	// that isn't necessarily encoded into the feed.
	Metadata string

	// Category is a thingfulx.Category instance - choose from the declared categories
	Category *Category

	// Webpage is a URL to a human intelligible page about the thing. This
	// doesn't have to be a unique page for the thing as typically one might not
	// be available. In that case, this could be a link to a page describing the
	// type of thing, or just a link to the main homepage of the project.
	Webpage string

	// DataURL is a string containing the actual URL we go to fetch data from for
	// this thing. This string must be unique for every thing indexed. We can
	// achieve this for things that don't make a real unique URL available by
	// adding a synthetic identifier to the URL string, i.e.
	// http://example.com?city=Indiana#section_id=1234
	DataURL string

	// IndexedAt is a Time object indicating the timestamp when the thing was
	// indexed by Thingful
	IndexedAt time.Time

	// Lng is the longitude of the Thing, represented as a decimal value
	Lng float64

	// Lat is the latitude of the Thing, represented as a decimal value
	Lat float64

	// Provider is a data structure containing information about the upstream
	// data provider.
	Provider *Provider

	// Visibility indicates what level of access the upstream data provider
	// offers, Closed, Shared or Open.
	Visibility Visibility

	// Channels are a slice of Channel objects which capture the actual data of
	// the thing
	Channels []Channel
}
