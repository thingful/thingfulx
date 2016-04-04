package thingfulx

import (
	"net/url"
)

// Provider is a data type used to hold information about a specific upstream data provider.
type Provider struct {
	// UID is a unique string identifying the provider. Good practice is to just use the package name
	UID string

	// Name is the name of the data provider
	Name string

	// Webpage should contain a link to the home page of the upstream provider.
	Webpage *url.URL

	// Description should contain some text describing the data provider.
	Description string
}
