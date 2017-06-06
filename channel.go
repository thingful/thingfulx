package thingfulx

import (
	"time"

	"github.com/thingful/thingfulx/schema"
)

// Channel represents an individual data channel being published by a Thing out
// there in the world somewhere.
type Channel struct {
	ID           string        `json:"id"`
	Metadata     []Metadata    `json:"metadata"`
	Observations []Observation `json:"observations"`

	// this is a bit confusing, can we use `json:"data_type"` instead?
	Type schema.DataType `json:"type"`

	// QuantityKind defines measurement type of quantity kind of this channel,
	// in Ontology PoV, Channel is type QuantityKind, so type of channel can be anything that classified as quantitykind
	// for example m3-lite:Temperature, m3-lite:WindSpeed
	QuantityKind string `json:"quantityKind"`

	// Unit defines unit of this channel, it is short for thingful:hasUnit,
	// for example m3-lite:PPM, m3-lite:Miles
	Unit string `json:"unit"`

	// DomainOfInterest defines domain of interest of this channel, it is short for thingful:hasDomainOfInterest,
	// for example m3-lite:Environment, m3-lite:Weather
	DomainOfInterest string `json:"domainOfInterest"`

	// MeasuredBy defines type of sensor that measure this channel, it is short for thingful:measuredBy,
	// for example m3-lite:Seismometer, m3-lite:LightSensor
	MeasuredBy string `json:"measuredBy"`
}

// Observation contains information about where and when the Observation
// was recorded.
type Observation struct {
	RecordedAt time.Time   `json:"recordedAt"`
	Location   *Location   `json:"location"`
	Val        interface{} `json:"value"`
}
