package thingfulx

import (
	"fmt"
	"time"
)

// Thing is the top level datastructure representing an individual thing as
// indexed on some upstream data provider.
type Thing struct {
	Title       string
	Description string
	Category    category
	Webpage     string
	URI         string
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

// String returns a String representation of the Thing (required for the
// Stringer interface)
func (t Thing) String() string {
	return fmt.Sprintf("Thing<Title:%s,URI:%s>", t.Title, t.URI)
}
