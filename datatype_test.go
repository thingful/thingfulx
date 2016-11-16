package thingfulx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataType(t *testing.T) {
	assert.Equal(t, "number", NumberType.String())
	assert.Equal(t, "string", StringType.String())
	assert.Equal(t, "boolean", BooleanType.String())
	assert.Equal(t, "dateTime", DateTimeType.String())
	assert.Equal(t, "url", URLType.String())
	assert.Equal(t, "binary", BinaryType.String())
}
