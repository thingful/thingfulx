package thingfulx

import (
	"fmt"
)

// Error is a type alias we use to create const possible error instances.
type Error string

// Error is the implementation of the error interface for our Error type.
func (e Error) Error() string { return string(e) }

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

// NewErrUnexpectedResponse is a constructor for creating ErrUnexpectedResponse
// instances
func NewErrUnexpectedResponse(status string) *ErrUnexpectedResponse {
	return &ErrUnexpectedResponse{status}
}

// ErrMissingConfig is a simple struct for reporting a missing config variable.
// It contains a single Variable attribute which should contain the name of the
// missing variable.
type ErrMissingConfig struct {
	Variable string
}

// Error is the implementation of the default error interface. Returns a simple
// string reporting the received status
func (e *ErrMissingConfig) Error() string {
	return fmt.Sprintf("thingfulx: missing config variable - '%s'", e.Variable)
}

// NewErrMissingConfig is a constructor for creating ErrMissingConfig instances
func NewErrMissingConfig(variable string) *ErrMissingConfig {
	return &ErrMissingConfig{variable}
}

// ErrBadData is an error that can be returned for any bad data when parsing
type ErrBadData struct {
	Msg string
}

// Error is the implementation of the default error interface. Returns the
// contained message
func (e *ErrBadData) Error() string {
	return fmt.Sprintf("thingfulx: %s", e.Msg)
}

// NewErrBadData is a constructor for creating ErrBadData instances
func NewErrBadData(msg string) *ErrBadData {
	return &ErrBadData{msg}
}

const (
	// ErrInvalidData is a generic error that can be returned if retrieved data
	// is invalid in any way
	ErrInvalidData = Error("thingfulx: invalid data")

	// ErrInvalidLocation is an error indicating that there is a problem with the
	// location of a fetched thing
	ErrInvalidLocation = Error("thingfulx: missing or invalid location data")

	// ErrInvalidTime is an error indicating that there is a problem with parsing
	// the time value from a retrieved thing
	ErrInvalidTime = Error("thingfulx: invalid time data")

	// ErrInvalidURL is an error used to flag an invalid url
	ErrInvalidURL = Error("thingfulx: invalid url")

	// ErrNotFound is an error that can be used for a 404 not found response
	ErrNotFound = Error("thingfulx: unexpected HTTP response code, got 404")
)
