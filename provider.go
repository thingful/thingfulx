package thingfulx

// Provider is a data structure containing information about the upstream
// data provider.
type Provider struct {
	// UID contains a namespaced URI for the provider, e.g.
	// thingful:openweathermap
	UID string `json:"id"`

	// Name contains a human friendly name of the provider
	Name string `json:"name"`

	// Description contains a human description of the provider
	Description string `json:"description,omitempty"`

	// Webpage should contain a link to a high level public web site about the
	// provider.
	Webpage string `json:"webpage"`
}
