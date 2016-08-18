package thingfulx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrUnexpectedResponse(t *testing.T) {
	e := NewErrUnexpectedResponse("404")

	msg := "thingfulx: unexpected HTTP response code, got 404"

	assert.Equal(t, e.Error(), msg)
}

func TestErrNotFound(t *testing.T) {
	msg := "thingfulx: unexpected HTTP response code, got 404"

	assert.Equal(t, ErrNotFound.Error(), msg)
}

func TestErrMissingConfig(t *testing.T) {
	e := NewErrMissingConfig("API_KEY")

	msg := "thingfulx: missing config variable - 'API_KEY'"

	assert.Equal(t, e.Error(), msg)
}

func TestErrBadData(t *testing.T) {
	e := NewErrBadData("some bad data")

	msg := "thingfulx: some bad data"

	assert.Equal(t, e.Error(), msg)
}
