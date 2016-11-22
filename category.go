package thingfulx

// Category is a type representing the high level category of a Thing. This
// implementation is kept for backwards compatability, ultimately this should
// be removed and folded into general metadata.
type Category struct {
	// Name is the name of the category
	Name string

	// UID is a unique URL safe ID
	UID string
}

var (
	// Home is an exported const representing the Thingful Home category
	Home = Category{
		Name: "Home",
		UID:  "home",
	}

	// Environment is an exported const representing the Thingful Environment category
	Environment = Category{
		Name: "Environment",
		UID:  "environment",
	}

	// Transport is an exported const representing the Thingful Transport category
	Transport = Category{
		Name: "Transport",
		UID:  "transport",
	}

	// Health is an exported const representing the Thingful Health category
	Health = Category{
		Name: "Health",
		UID:  "health",
	}

	// Energy is an exported const representing the Thingful Energy category
	Energy = Category{
		Name: "Energy",
		UID:  "energy",
	}

	// Flora is an exported const representing the Thingful Flora and Fauna category
	Flora = Category{
		Name: "Flora & Fauna",
		UID:  "flora-and-fauna",
	}

	// Experiment is an exported const representing the Thingful Experiment category
	Experiment = Category{
		Name: "Experiment",
		UID:  "experiment",
	}

	// Miscellaneous is an exported const representing the Thingful Miscellaneous category
	Miscellaneous = Category{
		Name: "Miscellaneous",
		UID:  "miscellaneous",
	}
)
