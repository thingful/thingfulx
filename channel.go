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
}

// Observation contains information about where and when the Observation
// was recorded.
type Observation struct {
	RecordedAt time.Time `json:"recorded_at"`
	Location   *Location `json:"location"`
	Data       *Data     `json:"data"`
}

// Data holds the actual value returned by the Thing as well as the
// concrete type of the value itself.
type Data struct {
	Type schema.DataType `json:"type"`
	Val  interface{}     `json:"value"`
}
