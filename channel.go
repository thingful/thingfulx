package thingfulx

import (
	"time"

	"github.com/thingful/thingfulx/schema"
)

// Channel represents an individual data channel being published by a Thing out
// there in the world somewhere.
type Channel struct {
	ID           string          `json:"id"`
	Metadata     []Metadata      `json:"metadata"`
	Observations []Observation   `json:"observations"`
	Type         schema.DataType `json:"type"`
}

// Observation contains information about where and when the Observation
// was recorded.
type Observation struct {
	RecordedAt time.Time   `json:"recorded_at"`
	Location   *Location   `json:"location"`
	Val        interface{} `json:"value"`
}
