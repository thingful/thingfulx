package thingfulx

import (
	"net/url"
)

// Provider is a data type used to hold information about a specific upstream data provider.
type Provider struct {
	// Name is the name of the data provider
	Name string

	// Webpage should contain a link to the home page of the upstream provider.
	Webpage *url.URL

	// DataURL should contain a URL representing the base URL data is actually
	// retrieved from. This might be different from the human intended Webpage
	// described previously
	DataURL *url.URL

	// Description should contain some text describing the data provider.
	Description string
}