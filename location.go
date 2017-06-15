package thingfulx

// Location is a data structure which holds an individual Thing or Observations
// geographic coordinates
type Location struct {
	Lng       float64 `json:"long"`
	Lat       float64 `json:"lat"`
	Elevation float64 `json:"elevation,omitempty"`
}
