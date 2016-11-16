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

func TestVerifyNumberType(t *testing.T) {
	validNumbers := []string{"0", "12.2", "-83.12"}
	for _, number := range validNumbers {
		got, err := VerifyNumberType(number)
		assert.Nil(t, err)
		assert.Equal(t, number, got)
	}

	invalidNumbers := []string{"", " ", "foo", "12.2.4"}
	for _, number := range invalidNumbers {
		_, err := VerifyNumberType(number)
		assert.Error(t, err)
	}
}

func TestVerifyBooleanType(t *testing.T) {
	validBooleans := []string{"0", "1", "True", "TRUE", "false", "F", "f"}
	for _, val := range validBooleans {
		got, err := VerifyBooleanType(val)
		assert.Nil(t, err)
		assert.Equal(t, val, got)
	}

	invalidBooleans := []string{"", " ", "foo", "0.0"}
	for _, val := range invalidBooleans {
		_, err := VerifyBooleanType(val)
		assert.Error(t, err)
	}
}
