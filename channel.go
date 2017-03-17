package thingfulx

import "time"

// Channel represents an individual data channel being published by a Thing out
// there in the world somewhere.
type Channel struct {
	ID           string            `json:"id"`
	Metadata     []ChannelMetadata `json:"metadata"`
	Observations []Observation     `json:"observations"`
	Data         Data              `json:"data"`
}

type ChannelMetadata struct {
	Prop string `json:"property"`
	Val  string `json:"value"`
}

type Observation struct {
	RecordedAt time.Time `json:"recorded_at"`
	Location   Location  `json:"location"`
}

type Data struct {
	Type string      `json:"type"`
	Val  interface{} `json:"value"`
}
