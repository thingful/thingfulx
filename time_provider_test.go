package thingfulx

import (
	"testing"
	"time"
)

func TestMockTimeProvider(t *testing.T) {
	ts := time.Now()

	provider := MockTimeProvider{ts}

	if provider.Now() != ts {
		t.Errorf("Unexpected response to Now(), expected '%s', got '%s'", ts, provider.Now())
	}

	newTs := time.Now()

	provider.Set(newTs)

	if provider.Now() == ts {
		t.Errorf("Unexpected response to Now(), expected '%s', got '%s'", ts, provider.Now())
	}

	if provider.Now() != newTs {
		t.Errorf("Unexpected response to Now(), expected '%s', got '%s'", newTs, provider.Now())
	}
}
