package thingfulx

// Visibility is a simple type alias for string used to represent an endpoint's
// visibility.
type Visibility string

// String is our implementation of Stringer for the Visibility type
func (v Visibility) String() string {
	return string(v)
}

const (
	// Closed is the exported const representing a private endpoint requiring
	// authentication to use.
	Closed = Visibility("closed")

	// Shared is the exported const representing an endpoint that requires
	// authentication but is free to use once credentials are supplied.
	Shared = Visibility("shared")

	// Open is the exported const representing an open endpoint, i.e. requiring
	// no authentication to use.
	Open = Visibility("open")
)
