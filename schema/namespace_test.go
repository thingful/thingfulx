package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpand(t *testing.T) {

	testcases := []struct {
		input    string
		expected string
	}{
		{
			input:    "thingful:foobar",
			expected: "https://thingful.github.io/schema#foobar",
		},
		{
			input:    "m3-lite:foobar",
			expected: "http://purl.org/iot/vocab/m3-lite#foobar",
		},
		{
			input:    "ssn:foobar",
			expected: "http://purl.oclc.org/NET/ssnx/ssn#foobar",
		},
		{
			input:    "iot-lite:foobar",
			expected: "http://purl.oclc.org/NET/UNIS/fiware/iot-lite#foobar",
		},
		{
			input:    "xsd:foobar",
			expected: "http://www.w3.org/2001/XMLSchema#foobar",
		},
		{
			input:    "unknown-namespace",
			expected: "unknown-namespace",
		},
		{
			input:    "qu:foobar",
			expected: "http://purl.org/NET/ssnx/qu/qu#foobar",
		},
		{
			input:    "https://thingful.github.io/schema#foobar",
			expected: "https://thingful.github.io/schema#foobar",
		},
	}

	for _, testcase := range testcases {
		got := Expand(testcase.input)
		assert.Equal(t, testcase.expected, got)
	}
}

func TestCompact(t *testing.T) {

	testcases := []struct {
		input    string
		expected string
	}{
		{
			input:    "https://thingful.github.io/schema#foobar",
			expected: "thingful:foobar",
		},
		{
			input:    "http://purl.org/iot/vocab/m3-lite#foobar",
			expected: "m3-lite:foobar",
		},
		{
			input:    "http://purl.oclc.org/NET/ssnx/ssn#foobar",
			expected: "ssn:foobar",
		},
		{
			input:    "http://purl.oclc.org/NET/UNIS/fiware/iot-lite#foobar",
			expected: "iot-lite:foobar",
		},
		{
			input:    "http://www.w3.org/2001/XMLSchema#foobar",
			expected: "xsd:foobar",
		},
		{
			input:    "http://schema.org/Organization",
			expected: "http://schema.org/Organization",
		},
		{
			input:    "http://purl.org/NET/ssnx/qu/qu#foobar",
			expected: "qu:foobar",
		},
		{
			input:    "m3-lite:foobar",
			expected: "m3-lite:foobar",
		},
	}

	for _, testcase := range testcases {
		got := Compact(testcase.input)
		assert.Equal(t, testcase.expected, got)
	}
}
