package thingfulx

// Metadata is a data structure containing additional metadata about a Channel
type Metadata struct {
	// Prop is an attribute used to hold the name of a property expressed as
	// either a namespaced string or a full URL.
	Prop string `json:"prop"`

	// Val is an attribute used to hold the value of the specified property.
	Val string `json:"val"`
}
