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
			expected: "http://purl.org/iot/vocab/thingful#foobar",
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
			input:    "schema:foobar",
			expected: "http://schema.org/foobar",
		},
		{
			input:    "jsonapi:foobar",
			expected: "http://purl.org/iot/vocab/jsonapi#foobar",
		},
		{
			input:    "thingfulqu:foobar1",
			expected: "http://purl.org/iot/vocab/thingful/qu#foobar1",
		},
		{
			input:    "cc:foobar",
			expected: "https://creativecommons.org/ns#foobar",
		},
		{
			input:    "biotop:foobar",
			expected: "http://purl.org/biotop/biotop.owl#foobar",
		},
		{
			input:    "datex:foobar",
			expected: "http://vocab.datex.org/terms#foobar",
		},
		{
			input:    "mobivoc:foobar",
			expected: "http://schema.mobivoc.org/foobar",
		},
		{
			input:    "siri:foobar",
			expected: "http://purl.org/iot/vocab/siri#foobar",
		},
		{
			input:    "http://purl.org/iot/vocab/thingful#foobar",
			expected: "http://purl.org/iot/vocab/thingful#foobar",
		},
		{
			input:    "sem:foobar",
			expected: "http://semanticweb.cs.vu.nl/2009/11/sem/foobar",
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
			input:    "http://purl.org/iot/vocab/thingful#foobar",
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
			expected: "schema:Organization",
		},
		{
			input:    "https://creativecommons.org/ns#license",
			expected: "cc:license",
		},
		{
			input:    "http://purl.org/NET/ssnx/qu/qu#foobar",
			expected: "qu:foobar",
		},
		{
			input:    "http://purl.org/iot/vocab/thingful/qu#foobar",
			expected: "thingfulqu:foobar",
		},
		{
			input:    "http://purl.org/iot/vocab/jsonapi#foobar",
			expected: "jsonapi:foobar",
		},
		{
			input:    "http://purl.org/biotop/biotop.owl#foobar",
			expected: "biotop:foobar",
		},
		{
			input:    "http://vocab.datex.org/terms#foobar",
			expected: "datex:foobar",
		},
		{
			input:    "http://schema.mobivoc.org/foobar",
			expected: "mobivoc:foobar",
		},
		{
			input:    "http://purl.org/iot/vocab/siri#foobar",
			expected: "siri:foobar",
		},
		{
			input:    "m3-lite:foobar",
			expected: "m3-lite:foobar",
		},
		{
			input:    "http://semanticweb.cs.vu.nl/2009/11/sem/foobar",
			expected: "sem:foobar",
		},
	}

	for _, testcase := range testcases {
		got := Compact(testcase.input)
		assert.Equal(t, testcase.expected, got)
	}
}
