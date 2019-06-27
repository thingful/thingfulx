package thingfulx

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClock(t *testing.T) {
	clock := NewClock()

	clock.Now()
}

func TestMockClockAt(t *testing.T) {
	ts := time.Now()

	clock := NewMockClockAt(ts)

	assert.Equal(t, clock.Now(), ts)
}

func TestMockClock(t *testing.T) {
	clock := NewMockClock()
	ts := clock.Now()
	assert.NotNil(t, ts)
	assert.Equal(t, ts, clock.Now())
}
