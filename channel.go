package thingfulx

import (
	"time"
)

// Channel represents an individual data channel being published by a Thing out
// there in the world somewhere.
type Channel struct {
	ID         string
	Value      string
	RecordedAt time.Time
	Unit       Unit
}
