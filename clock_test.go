package thingfulx

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClock(t *testing.T) {
	provider := NewClock()

	provider.Now()
}

func TestMockClock(t *testing.T) {
	ts := time.Now()

	provider := NewMockClock(ts)

	assert.Equal(t, provider.Now(), ts)
}
