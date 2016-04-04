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

// NewTimeProvider returns an instantiated DefaultTimeProvider instance.
func NewTimeProvider() TimeProvider {
	return &DefaultTimeProvider{}
}

// DefaultTimeProvider is an implementation of our TimeProvider interface that
// returns now directly from time.Now().
type DefaultTimeProvider struct{}

// Now returns the current time. In the default implementation this is just the
// value of time.Now()
func (t DefaultTimeProvider) Now() time.Time {
	return time.Now()
}

// NewMockTimeProvider is a convenience helper that returns a new
// MockTimeProvider initialized to the specified time.
func NewMockTimeProvider(t time.Time) TimeProvider {
	return &MockTimeProvider{
		InternalTime: t,
	}
}

// MockTimeProvider is an implementation of the TimeProvider interface that
// gives back a canned time for all calls to Now(). Intended for use in testing
// Fetchers.
type MockTimeProvider struct {
	InternalTime time.Time
}

// Now returns the internal time of the MockTimeProvider.
func (t MockTimeProvider) Now() time.Time {
	return t.InternalTime
}
