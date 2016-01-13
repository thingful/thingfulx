package thingfulx

import (
	"errors"
	"fmt"
)

// ErrUnexpectedResponse is a simple struct for reporting unexpected HTTP
// responses. It contains a single attribute which allows capturing the
// received HTTP status
type ErrUnexpectedResponse struct {
	Status string
}

// Error is the implementation of the default error interface. Returns a simple
// string reporting the received status
func (e *ErrUnexpectedResponse) Error() string {
	return fmt.Sprintf("thingfulx: unexpected HTTP response code, got %s", e.Status)
}

var (
	// ErrInvalidData is a generic error that can be returned if retrieved data
	// is invalid in any way
	ErrInvalidData = errors.New("thingfulx: invalid data")

	// ErrInvalidLocationData is an error indicating that there is a problem with
	// the location of a fetched thing
	ErrInvalidLocationData = errors.New("thingfulx: missing or invalid location data")

	// ErrInvalidTimeData is an error indicating that there is a problem with
	// parsing the time value from a retrieved thing
	ErrInvalidTimeData = errors.New("thingfulx: invalid time data")

	// ErrMissingCredentials is an error that should be returned if a fetcher is
	// missing required credentials
	ErrMissingCredentials = errors.New("thingfulx: missing required credentials")
)
