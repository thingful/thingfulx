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
	IndexedAt time.Time `json:"indexedAt"`

	// Webpage is a URL to a human intelligible page about the thing. This
	// doesn't have to be a unique page for the thing as typically one might not
	// be available. In that case, this could be a link to a page describing the
	// type of thing, or just a link to the main homepage of the project.
	Webpage string `json:"webpage"`

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

	// Visibility indicates generally what type of access is possible to the
	// Thing using the data spectrum classifications from the ODI: Open, Shared
	// or Closed
	Visibility Visibility `json:"visibility"`

	// ThingType indicates what type is this thing, it is short for thingful:isThingType
	// the most generic thing type is thingful:ConnectedDevice
	ThingType string `json:"thingType"`

	// Category indicates what type this thing belongs to, it is short for
	// thingful:hasCategory This is mainly used by thingful.net.  The value will
	// be those defined by m3-lite's domain of interest, plus some additionally
	// defined by thingful
	Category Category `json:"category"`

	// DataLicense indicates what of data license this thing has, it is short for
	// cc:license This could be used to decide what can we do with this thing's
	// data It takes DataLicense struct which can be manually defined or use
	// predefined ones
	DataLicense *DataLicense `json:"license"`

	// AttributionName indicates name ot the author of the data (if required)
	// it is short for cc:attributionName
	AttributionName string `json:"attributionName,omitempty"`

	// AttributionURL indicates name ot the author of the data (if required)
	// it is short for cc:attributionURL
	AttributionURL string `json:"attributionURL,omitempty"`

	// UpdateInterval indicates how often this resource should get updated (in
	// seconds)
	UpdateInterval int `json:"updateInterval,omitempty"`
}

// Provider is a data structure containing information about the upstream
// data provider.
type Provider struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url"`
}

// Endpoint is a data structure containing the actual URL we go to fetch data
// from for this thing, the page MIME type and information about the endpoint
// authentication. The URL value must be unique for every thing indexed.
// We can achieve this for things that don't make a real unique URL available
// by adding a synthetic identifier to the URL string, i.e.
// http://example.com?city=Indiana#section_id=1234
type Endpoint struct {
	URL         string `json:"url"`
	ContentType string `json:"contentType"`
}

// MetaVal returns the first value for the given property in the Thing's
// metadata. If the property is not found this returns the empty string. If you
// need to access multiple values with the same property key, you should use
// the map directly.
func (t *Thing) MetaVal(prop string) string {
	for _, meta := range t.Metadata {
		if meta.Prop == prop {
			return meta.Val
		}
	}

	return ""
}
