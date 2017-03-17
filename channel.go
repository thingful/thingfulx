package thingfulx

import "time"

// Channel represents an individual data channel being published by a Thing out
// there in the world somewhere.
type Channel struct {
	ID           string            `json:"id"`
	Metadata     []ChannelMetadata `json:"metadata"`
	Observations []Observation     `json:"observations"`
}

// ChannelMetadata is an object which contains any metadata required to
// semantically descibe a measurement collected by a Thing.
type ChannelMetadata struct {
	Prop string `json:"property"`
	Val  string `json:"value"`
}

// Observation contains information about where and when the Observation
// was recorded.
type Observation struct {
	RecordedAt time.Time `json:"recorded_at"`
	Location   Location  `json:"location"`
	Data       Data      `json:"data"`
}

// Data holds the actual value returned by the Thing as well as the
// concrete type of the value itself.
type Data struct {
	Type string      `json:"type"`
	Val  interface{} `json:"value"`
}
