package thingfulx

import (
	"time"
)

// Thing is the top level datastructure representing an individual thing as
// indexed on some upstream data provider.
type Thing struct {
	// Title is a short descriptive lable for the thing
	Title string `json:"title"`

	// Description is a long form description about the thing
	Description string `json:"description"`

	// Metadata is a comma separated tags - additional metadata about the Thing,
	// that isn't necessarily encoded into the feed.
	Metadata string `json:"-"`

	// Category is a thingfulx.Category instance - choose from the declared categories
	Category *Category `json:"-"`

	// Webpage is a URL to a human intelligible page about the thing. This
	// doesn't have to be a unique page for the thing as typically one might not
	// be available. In that case, this could be a link to a page describing the
	// type of thing, or just a link to the main homepage of the project.
	Webpage string `json:"webpage"`

	// DataURL is a string containing the actual URL we go to fetch data from for
	// this thing. This string must be unique for every thing indexed. We can
	// achieve this for things that don't make a real unique URL available by
	// adding a synthetic identifier to the URL string, i.e.
	// http://example.com?city=Indiana#section_id=1234
	DataURL string `json:"datasource"`

	// IndexedAt is a Time object indicating the timestamp when the thing was
	// indexed by Thingful
	IndexedAt time.Time `json:"indexed_at"`

	// Lng is the longitude of the Thing, represented as a decimal value
	Lng float64 `json:"longitude"`

	// Lat is the latitude of the Thing, represented as a decimal value
	Lat float64 `json:"latitude"`

	// Provider is a data structure containing information about the upstream
	// data provider.
	Provider *Provider `json:"-"`

	// Visibility indicates what level of access the upstream data provider
	// offers, Closed, Shared or Open.
	Visibility Visibility `json:"visibility"`

	// RawData is the original data returned by the upstream data provider
	RawData interface{} `json:"raw_data"`

	// Channels are a slice of Channel objects which capture the actual data of
	// the thing
	Channels []Channel `json:"channels"`
}
