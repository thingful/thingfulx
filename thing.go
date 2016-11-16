package thingfulx

import (
	"net/url"
	"time"
)

// Thing is the top level datastructure representing an individual thing as
// indexed on some upstream data provider.
type Thing struct {
	// Title is a short descriptive lable for the thing
	Title string

	// Description is a long form description about the thing
	Description string

	// Metadata is a slice of Tags used to describe some features of the Thing
	Metadata []Tag

	// Webpage is a URL to a human intelligible page about the thing. This
	// doesn't have to be a unique page for the thing as typically one might not
	// be available. In that case, this could be a link to a page describing the
	// type of thing, or just a link to the main homepage of the project.
	Webpage *url.URL

	// DataURL is a string containing the actual URL we go to fetch data from for
	// this thing. This string must be unique for every thing indexed. We can
	// achieve this for things that don't make a real unique URL available by
	// adding a synthetic identifier to the URL string, i.e.
	// http://example.com?city=Indiana#section_id=1234
	DataURL *url.URL

	// IndexedAt is a Time object indicating the timestamp when the thing was
	// indexed by Thingful
	IndexedAt time.Time

	// Location is a sturct used to hold a location of the Thing
	Location *Location

	// Provider is a data structure containing information about the upstream
	// data provider.
	Provider *Provider

	// RawData is the original data returned by the upstream data provider
	RawData []byte

	// Channels are a slice of Channel objects which capture the actual data of
	// the thing
	Channels []Channel
}
