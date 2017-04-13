package thingfulx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	assert.Equal(t, "closed", Closed.String())
}
