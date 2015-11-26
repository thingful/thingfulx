package thingfulx

// Visibility is a simple type alias for string.
type Visibility string

const (
	// Closed is the exported const representing a private Thing
	Closed = Visibility("closed")

	// Shared is the exported const representing a Thing accessible through authentication
	Shared = Visibility("shared")

	// Open is the exported const representing an open Thing
	Open = Visibility("open")
)
