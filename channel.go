package thingfulx

import "time"

// Channel represents an individual data channel being published by a Thing out
// there in the world somewhere.
type Channel struct {
	ID           string        `json:"id"`
	Metadata     []Metadata    `json:"metadata"`
	Observations []observation `json:"observations"`
	Data         data          `json:"data"`
}

type observation struct {
	RecordedAt time.Time `json:"recorded_at"`
	Location   Location  `json:"location"`
}

type data struct {
	Type string      `json:"type"`
	Val  interface{} `json:"value"`
}
