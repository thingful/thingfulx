package thingfulx

import (
	"time"
)

// Channel represents an individual data channel being published by a Thing out
// there in the world somewhere.
type Channel struct {
	ID         string    `json:"id"`
	Value      string    `json:"value"`
	RecordedAt time.Time `json:"recorded_at"`
	Unit       string    `json:"units"`
}
