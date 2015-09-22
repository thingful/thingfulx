package thingfulx

// Parser is the interface for things that know how to turn a raw byte array
// retrieved from some upstream source, and turn it into a slice of Things
// ready to be persisted on the Thingful side.
type Parser interface {

	// The main Parse function - it takes in a slice of bytes that will have been
	// retrieved from somewhere, and transforms them into a Slice of Thing
	// instances.
	Parse(data []byte) ([]Thing, error)
}
