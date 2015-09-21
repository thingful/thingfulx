package thingfulx

type category string

const (
	// Home is the exported const representing the Thingful Home category
	Home = category("home")

	// Environment is the exported const representing the Thingful Environment category
	Environment = category("environment")

	// Transport is the exported const representing the Thingful Transport category
	Transport = category("transport")

	// Health is the exported const representing the Thingful Health category
	Health = category("health")

	// Energy is the exported const representing the Thingful Energy category
	Energy = category("energy")

	// Flora is the exported const representing the Thingful Flora category
	Flora = category("flora & fauna")

	// Experiment is the exported const representing the Thingful Experiment category
	Experiment = category("experiment")

	// Miscellaneous is the exported const representing the Thingful Miscellaneous category
	Miscellaneous = category("miscellaneous")
)
