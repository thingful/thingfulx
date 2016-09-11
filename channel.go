package thingfulx

import (
	"time"
)

// Channel represents an individual data channel being published by a Thing out
// there in the world somewhere.
type Channel struct {
	// Label is some arbitrary string label for the channel. A thing may contain
	// Channels with duplicated labels, but this makes parsing the data more
	// challenging
	Label string

	// Value the value of the channel at time of recording. This is an opaque
	// string value which can contain anything.
	Value string

	// RecordedAt encodes the timestamp at which the channel was recorded.
	RecordedAt time.Time

	// Sensitive is a boolean flag used to indicate that a specific channel
	// contains data the data owner regards as sensitive or that whoever is
	// writing the Fetcher regards as sensitive.
	Sensitive bool

	// Metadata contains a slice of Tag objects containing arbitrary additional
	// facts about the Channel that Thingful should index.
	Metadata []Tag
}
