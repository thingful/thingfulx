package thingfulx

import (
	"time"
)

// Thing is the top level data structure representing an individual thing as
// indexed on some upstream data provider.
type Thing struct {
	// Title is a short descriptive label for the thing
	Title string `json:"title"`

	// Description is a long form description about the thing
	Description string `json:"description"`

	// IndexedAt is a Time object indicating the timestamp when the thing was
	// indexed by Thingful
	IndexedAt time.Time `json:"indexed_at"`

	// Webpage is a URL to a human intelligible page about the thing. This
	// doesn't have to be a unique page for the thing as typically one might not
	// be available. In that case, this could be a link to a page describing the
	// type of thing, or just a link to the main homepage of the project.
	Webpage string `json:"-"`

	// DataURL is an embedded data structure which describes the actual data entry point.
	Endpoint *Endpoint `json:"endpoint"`

	// Metadata is a data structure containing additional metadata about the Thing.
	Metadata []Metadata `json:"metadata"`

	// Location is an embedded data structure which describes location data.
	Location *Location `json:"location"`

	// Provider is an embedded data structure which describes the data provider.
	Provider *Provider `json:"provider"`

	// Channels are a slice of Channel objects which capture the actual data of
	// the thing
	Channels []Channel `json:"channels"`
}

// Provider is a data structure containing information about the upstream
// data provider.
type Provider struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// Endpoint is a data structure containing the actual URL we go to fetch data
// from for this thing, the page MIME type and information about the endpoint
// authentication. The URL value must be unique for every thing indexed.
// We can achieve this for things that don't make a real unique URL available
// by adding a synthetic identifier to the URL string, i.e.
// http://example.com?city=Indiana#section_id=1234
type Endpoint struct {
	URL            string `json:"url"`
	ContentType    string `json:"content_type"`
	Authentication string `json:"authentication"`
}
