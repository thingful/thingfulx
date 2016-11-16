package thingfulx

import (
	"fmt"
	"strconv"
)

// Longitude is a type alias for float64 to try and eliminate errors with
// getting Long/Lat mixed up.
type Longitude float64

// Latitude is a type alias for float64 to try and eliminate errors with
// getting Long/Lat mixed up.
type Latitude float64

// Location is a struct used to hold location information about a Thing
type Location struct {
	// Longitude is the longitude of the Thing, represented as a decimal value
	Longitude Longitude

	// Latitude is the latitude of the Thing, represented as a decimal value
	Latitude Latitude
}

// NewLocation constructs a new Location object with the given longitude and
// Location values. However inside here we truncate spurious precision from the
// long/lat values:
//
// On the precision of decimal long/lat values:
//
// > The sixth decimal place is worth up to 0.11 m: you can use this for laying
// > out structures in detail, for designing landscapes, building roads. It
// > should be more than good enough for tracking movements of glaciers and
// > rivers. This can be achieved by taking painstaking measures with GPS, such
// > as differentially corrected GPS.
func NewLocation(lng Longitude, lat Latitude) *Location {
	return &Location{
		Longitude: Longitude(Truncate(float64(lng))),
		Latitude:  Latitude(Truncate(float64(lat))),
	}
}

// Truncate uses strconv's ParseFloat and fmt.Sprintf to truncate the given
// value to no more than 6 digits after the decimal point. For long/lat values
// this is as accurate as we need.
func Truncate(val float64) float64 {
	// I think we can safely ignore the error here, as we pass in a valid
	// float64, and all that happens here is that we truncate it's decimal
	// places.
	truncated, _ := strconv.ParseFloat(fmt.Sprintf("%.6f", val), 64)
	return truncated
}
