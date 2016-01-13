package thingfulx

import (
	"testing"
)

func TestErrUnexpectedResponse(t *testing.T) {
	e := ErrUnexpectedResponse{
		Status: "404",
	}

	msg := "thingfulx: unexpected HTTP response code, got 404"

	if e.Error() != msg {
		t.Errorf("Unexpected response, expected '%s', got '%s'", msg, e.Error())
	}
}
