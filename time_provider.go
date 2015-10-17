package thingfulx

import (
	"time"
)

// TimeProvider is an interface that allows us to test time dependent code
// easily by providing an implementation of this interface that returns a
// 'canned' time rather than the real time.
type TimeProvider interface {
	// Now returns the current time as seen by the provider. This could be a fake
	// time for testing time providers.
	Now() time.Time
}

// DefaultTimeProvider is an implementation of our TimeProvider interface that
// returns now directly from time.Now().
type DefaultTimeProvider struct{}

// Now just returns the time as seen by time.Now()
func (t *DefaultTimeProvider) Now() time.Time {
	return time.Now()
}

// MockTimeProvider is an implementation of the TimeProvider interface that
// gives back a canned time for all calls to Now(). Intended for use in testing
// Fetchers.
type MockTimeProvider struct {
	InternalTime time.Time
}

// Now returns the internal time of the MockTimeProvider.
func (t *MockTimeProvider) Now() time.Time {
	return t.InternalTime
}

// Set is a function that allows setting the internal time of the MockTimeProvider
func (t *MockTimeProvider) Set(other time.Time) {
	t.InternalTime = other
}
