package thingfulx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTruncate(t *testing.T) {
	testcases := []struct {
		input    float64
		expected float64
	}{
		{
			input:    0,
			expected: 0,
		},
		{
			input:    5.123456789,
			expected: 5.123457,
		},
		{
			input:    -180.987654321,
			expected: -180.987654,
		},
	}

	var got float64

	for _, testcase := range testcases {
		got = Truncate(testcase.input)
		assert.Equal(t, testcase.expected, got)
	}
}

func TestNewLocation(t *testing.T) {
	testcases := []struct {
		lng      Longitude
		lat      Latitude
		expected *Location
	}{
		{
			lng: 0,
			lat: 0,
			expected: &Location{
				Longitude: 0,
				Latitude:  0,
			},
		},
	}

	var got *Location

	for _, testcase := range testcases {
		got = NewLocation(testcase.lng, testcase.lat)
		assert.Equal(t, testcase.expected, got)
	}
}
