package thingfulx

import (
	"fmt"
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

// String returns a string representation of the Channel - required to
// implement the fmt.Stringer interface
func (c Channel) String() string {
	return fmt.Sprintf("Channel<ID:%s,Value:%s>", c.ID, c.Value)
}
