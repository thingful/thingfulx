package thingfulx

import (
	"time"

	"github.com/thingful/thingfulx/schema"
)

// Channel represents an individual data channel being published by a Thing out
// there in the world somewhere.
type Channel struct {
	// URL is a unique identifier for the channel expressed as a fully qualified
	// URL that must be the URL at which data is available for the Channel.
	URL string `json:"url"`

	// Title is a human centered title for the channel that ideally also uniquely
	// identifies the Channel amongst other Channels published by a Device
	Title string `json:"title"`

	// Description is a human centered description of the channel.
	Description string `json:"description"`

	// Webpage is a link to a web resource containing information about the
	// channel or device that is publishing this data.
	Webpage string `json:"webpage"`

	// IndexedAt is a Time object indicating the timestamp when the channel was
	// indexed by Thingful
	IndexedAt time.Time `json:"indexedAt"`

	// Location is an embedded data structure containing the physical location of
	// the channel at time of indexing.
	Location *Location `json:"location"`

	// Provider is an embedded data structure containing information about the
	// data provider.
	Provider *Provider `json:"provider"`

	// Metadata is a data structure that holds a list of Metadata properties about
	// the Channel. This is for facts that do not fit neatly into the primary
	// attributes defined on the channel.
	Metadata []Metadata `json:"metadata"`

	// Visibility indicates what type of access is possible to the data provider
	// using the data spectrum classifications developed by the ODI, one of Open,
	// Shared or Closed.
	Visibility Visibility `json:"visibility"`

	// DataLicense contains information about the licensing requirements on the
	// data as set by the provider.
	DataLicense *DataLicense `json:"dataLicense"`

	// AttributionName is used to describe an optional attribution name that the
	// provider may require to be included with any use of their data.
	AttributionName string `json:"attributionName,omitempty"`

	// AttributionURL is used to describe an optional attribution URL that the
	// provider may require to be included with any use of their data.
	AttributionURL string `json:"attributionURL,omitempty"`

	// UpdateInterval indicates how often the upstream resource obtains fresh data
	// expressed as an integer number of seconds.
	UpdateInterval int32 `json:"updateInterval,omitempty"`

	// Type defines the type of data of this channel, eg; string, double,
	// date-time
	Type schema.DataType `json:"dataType"`

	// ContentType is used to encode the mime content type of the published data.
	ContentType string `json:"contentType,omitempty"`

	// QuantityKind defines measurement type of quantity kind of this channel, in
	// Ontology PoV, Channel is type QuantityKind, so type of channel can be
	// anything that classified as quantitykind for example m3-lite:Temperature,
	// m3-lite:WindSpeed
	QuantityKind string `json:"quantityKind"`

	// Unit defines unit of this channel, it is short for thingful:hasUnit, for
	// example m3-lite:PPM, m3-lite:Miles
	Unit string `json:"unit"`

	// DomainOfInterest defines domain of interest of this channel, it is short
	// for thingful:hasDomainOfInterest, for example m3-lite:Environment,
	// m3-lite:Weather
	DomainOfInterest []string `json:"domain"`

	// MeasuredBy defines type of sensor that measure this channel, it is short
	// for thingful:measuredBy, for example m3-lite:Seismometer,
	// m3-lite:LightSensor
	MeasuredBy string `json:"measuredBy"`

	// Observations contains a list of observations containing specific datapoints
	// for the Channel
	Observations []Observation `json:"observations"`
}

// Observation contains information about where and when the Observation
// was recorded.
type Observation struct {
	// RecordedAt is the timestamp at which the datapoint was recorded
	RecordedAt time.Time `json:"recordedAt"`

	// Location is the geographical location where the individual datapoint was
	// recorded. This is included separately from the main Location as a channel
	// may be mobile
	Location *Location `json:"location"`

	// Val contains the specific value of the datapoint expressed as a string.
	Val string `json:"value"`
}
