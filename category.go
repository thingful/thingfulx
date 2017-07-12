package thingfulx

// Category is our type alias for thingful category instances
type Category string

// String is our implementation of Stringer for the Category type
func (c Category) String() string {
	return string(c)
}

const (
	// Home is the exported const representing the Thingful Home category
	Home = Category("http://purl.org/iot/vocab/thingful#Home")

	// Environment is the exported const representing the Thingful Environment
	// category
	Environment = Category("http://purl.org/iot/vocab/thingful#Environment")

	// Transport is the exported const representing the Thingful Transport
	// category
	Transport = Category("http://purl.org/iot/vocab/thingful#Transport")

	// Health is the exported const representing the Thingful Health category
	Health = Category("http://purl.org/iot/vocab/thingful#Health")

	// Energy is the exported const representing the Thingful Energy category
	Energy = Category("http://purl.org/iot/vocab/thingful#Energy")

	// Flora is the exported const representing the Thingful Flora and Fauna
	// category
	Flora = Category("http://purl.org/iot/vocab/thingful#FloraFauna")

	// Experiment is the exported const representing the Thingful Experiment
	// category
	Experiment = Category("http://purl.org/iot/vocab/thingful#Experiment")

	// Miscellaneous is the exported const representing the Thingful
	// Miscellaneous category
	Miscellaneous = Category("http://purl.org/iot/vocab/thingful#Miscellaneous")
)
