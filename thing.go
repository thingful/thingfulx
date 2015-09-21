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

func (c Channel) String() string {
	return fmt.Sprintf("%s: %s - %s, %s", c.ID, c.Value, c.Unit, c.RecordedAt)
}

type Thing struct {
	Title       string
	Description string
	Category    category
	Datasource  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IndexedAt   time.Time
	Lng         float64
	Lat         float64
	Provider    string
	Visibility  string
	DataURL     string
	Channels    []Channel
}
