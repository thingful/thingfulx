package thingfulx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategoryString(t *testing.T) {
	assert.Equal(t, "http://purl.org/iot/vocab/thingful#Home", Home.String())
}
