package thingfulx

// Location is a data structure which holds an individual Thing or Observations
// geographic coordinates
type Location struct {
	Lng       float32 `json:"long"`
	Lat       float32 `json:"lat"`
	Elevation float32 `json:"elevation,omitempty"`
}
