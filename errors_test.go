package thingfulx

import (
	"testing"
)

func TestErrUnexpectedResponse(t *testing.T) {
	e := NewErrUnexpectedResponse("404")

	msg := "thingfulx: unexpected HTTP response code, got 404"

	if e.Error() != msg {
		t.Errorf("Unexpected response, expected '%s', got '%s'", msg, e.Error())
	}
}

func TestErrMissingConfig(t *testing.T) {
	e := NewErrMissingConfig("API_KEY")

	msg := "thingfulx: missing config variable - 'API_KEY'"

	if e.Error() != msg {
		t.Errorf("Unexpected response, expected '%s', got '%s'", msg, e.Error())
	}
}

func TestErrBadData(t *testing.T) {
	e := NewErrBadData("some bad data")

	msg := "thingfulx: some bad data"

	if e.Error() != msg {
		t.Errorf("Unexpected response, expected '%s', got '%s'", msg, e.Error())
	}
}
