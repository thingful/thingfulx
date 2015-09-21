package thingfulx

import (
	"fmt"
)

// Unit is a struct representing a unit (e.g. grams, ohms or sieverts)
type Unit struct {
	Name   string
	Symbol string
}

// String is the implementation of the fmt.Stringer interface. This method is
// used to print out a representation of the type.
func (u Unit) String() string {
	return fmt.Sprintf("%s (%s)", u.Name, u.Symbol)
}
