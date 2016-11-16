package thingfulx

import (
	"time"
)

// Channel represents an individual data channel being published by a Thing out
// there in the world somewhere.
type Channel struct {
	Value      string
	Type       DataType
	RecordedAt time.Time
	Metadata   []Tag
}
