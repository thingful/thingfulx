package thingfulx

import (
	"testing"
	"time"
)

func TestTimeProvider(t *testing.T) {
	provider := NewTimeProvider()

	provider.Now()
}

func TestMockTimeProvider(t *testing.T) {
	ts := time.Now()

	provider := NewMockTimeProvider(ts)

	if provider.Now() != ts {
		t.Errorf("Unexpected response to Now(), expected '%s', got '%s'", ts, provider.Now())
	}
}
