package thingfulx

// Category is a simple type alias for a string
type Category struct {
	// Name is the name of the category
	Name string

	// UID is a unique URL safe ID
	UID string
}

var (
	// Home is the exported const representing the Thingful Home category
	Home = Category{
		Name: "Home",
		UID:  "home",
	}

	// Environment is the exported const representing the Thingful Environment category
	Environment = Category{
		Name: "Environment",
		UID:  "environment",
	}

	// Transport is the exported const representing the Thingful Transport category
	Transport = Category{
		Name: "Transport",
		UID:  "transport",
	}

	// Health is the exported const representing the Thingful Health category
	Health = Category{
		Name: "Health",
		UID:  "health",
	}

	// Energy is the exported const representing the Thingful Energy category
	Energy = Category{
		Name: "Energy",
		UID:  "energy",
	}

	// Flora is the exported const representing the Thingful Flora and Fauna category
	Flora = Category{
		Name: "Flora & Fauna",
		UID:  "flora-and-fauna",
	}

	// Experiment is the exported const representing the Thingful Experiment category
	Experiment = Category{
		Name: "Experiment",
		UID:  "experiment",
	}

	// Miscellaneous is the exported const representing the Thingful Miscellaneous category
	Miscellaneous = Category{
		Name: "Miscellaneous",
		UID:  "miscellaneous",
	}
)
