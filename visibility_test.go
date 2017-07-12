package thingfulx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVisibilityString(t *testing.T) {
	assert.Equal(t, "closed", Closed.String())
}
