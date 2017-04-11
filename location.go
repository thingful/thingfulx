package thingfulx

// Location is a data structure which holds an individual Thing or Observations
// geographic coordinates
type Location struct {
	Lng       float64 `json:"longitude"`
	Lat       float64 `json:"latitude"`
	Elevation float64 `json:"elevation,omitempty"`
}
