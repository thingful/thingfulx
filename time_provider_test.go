package thingfulx

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeProvider(t *testing.T) {
	provider := NewTimeProvider()

	provider.Now()
}

func TestMockTimeProvider(t *testing.T) {
	ts := time.Now()

	provider := NewMockTimeProvider(ts)

	assert.Equal(t, provider.Now(), ts)
}
