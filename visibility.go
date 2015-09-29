package thingfulx

type visibility string

const (
	// Public is the exported const representing a Public thing
	Public = visibility("public")

	// Private is the exported const representing a Private thing
	Private = visibility("private")
)
