package thingfulx

import (
	"testing"
)

func TestThingStringer(t *testing.T) {
	expected := "Thing<Title:Weather,URI:http://example.com/pages/foo>"

	thing := Thing{
		Title: "Weather",
		URI:   "http://example.com/pages/foo",
	}

	if thing.String() != expected {
		t.Errorf("Unexpected String() output, expected '%s', got '%s'", expected, thing.String())
	}
}
