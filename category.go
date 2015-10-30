package thingfulx

type Category string

const (
	// Home is the exported const representing the Thingful Home category
	Home = Category("home")

	// Environment is the exported const representing the Thingful Environment category
	Environment = Category("environment")

	// Transport is the exported const representing the Thingful Transport category
	Transport = Category("transport")

	// Health is the exported const representing the Thingful Health category
	Health = Category("health")

	// Energy is the exported const representing the Thingful Energy category
	Energy = Category("energy")

	// Flora is the exported const representing the Thingful Flora category
	Flora = Category("flora & fauna")

	// Experiment is the exported const representing the Thingful Experiment category
	Experiment = Category("experiment")

	// Miscellaneous is the exported const representing the Thingful Miscellaneous category
	Miscellaneous = Category("miscellaneous")
)
