package thingfulx

type Visibility string

const (
	// Public is the exported const representing a Public thing
	Public = Visibility("public")

	// Private is the exported const representing a Private thing
	Private = Visibility("private")
)
