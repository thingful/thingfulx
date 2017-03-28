package thingfulx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpandNamespace(t *testing.T) {

	testcases := []struct {
		input    string
		expected Ontology
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
			expected: "",
		},
	}

	for _, testcase := range testcases {
		got := ExpandNamespace(testcase.input)

		assert.Equal(t, testcase.expected, got)
	}
}
