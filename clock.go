package thingfulx

import (
	"time"
)

// Clock is an interface that allows us to test time dependent code
// easily by providing an implementation of this interface that returns a
// 'canned' time rather than the real time.
type Clock interface {
	// Now returns the current time as seen by the provider. This could be a fake
	// time for testing time providers.
	Now() time.Time
}

// TimeProvider is an alias for our Clock interface for backwards compatibility
type TimeProvider = Clock

// NewClock returns an instantiated DefaultClock instance.
func NewClock() Clock {
	return &DefaultClock{}
}

// DefaultClock is an implementation of our TimeProvider interface that
// returns now directly from time.Now().
type DefaultClock struct{}

// Now returns the current time. In the default implementation this is just the
// value of time.Now()
func (c DefaultClock) Now() time.Time {
	return time.Now()
}

// NewMockClock returns a new MockClock with the internal time set to the time
// at which the method was invoked.
func NewMockClock() Clock {
	return &mockClock{
		internalTime: time.Now(),
	}
}

// NewMockClockAt returns a new MockClock initialized to the passed in time
// variable.
func NewMockClockAt(t time.Time) Clock {
	return &mockClock{
		internalTime: t,
	}
}

// mockClock is an implementation of the Clock interface that gives back a
// canned time for all calls to Now(). Intended for use in testing Fetchers.
type mockClock struct {
	internalTime time.Time
}

// Now returns the internal time of the MockTimeProvider.
func (t mockClock) Now() time.Time {
	return t.internalTime
}

// NewMockTimeProvider returns a mock time provider suitable for use in tests.
//
// Deprecated: the TimeProvider type has been renamed to Clock, with matching
// mock instance constructors. New code should use the Clock versions.
func NewMockTimeProvider(t time.Time) Clock {
	return NewMockClockAt(t)
}
