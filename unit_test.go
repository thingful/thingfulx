package thingfulx

import (
	"testing"
)

func TestUnitStringer(t *testing.T) {
	u := Unit{
		Name:   "gram",
		Symbol: "g",
	}

	if u.String() != "Unit<Name:gram,Symbol:g>" {
		t.Errorf("Unexpected String output, expected 'Unit<gram,g>', got '%s'", u.String())
	}
}
