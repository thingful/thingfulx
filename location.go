package thingfulx

// Location is a data structure which holds an individual Thing's
// geographic coordinates
type Location struct {
	Lng float64 `json:"longitude"`
	Lat float64 `json:"latitude"`
}
