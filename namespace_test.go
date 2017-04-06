package thingfulx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpand(t *testing.T) {

	testcases := []struct {
		input    string
		expected string
		err      error
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
			err:      ErrUnknownNamespace,
		},
	}

	for _, testcase := range testcases {
		got, err := Expand(testcase.input)
		if testcase.err == nil {
			assert.Nil(t, err)
			assert.Equal(t, testcase.expected, got)
		} else {
			assert.Equal(t, testcase.err, err)
		}
	}
}

func TestCompact(t *testing.T) {

	testcases := []struct {
		input    string
		expected string
		err      error
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
			expected: "",
			err:      ErrUnknownNamespace,
		},
	}

	for _, testcase := range testcases {
		got, err := Compact(testcase.input)
		if testcase.err == nil {
			assert.Nil(t, err)
			assert.Equal(t, testcase.expected, got)
		} else {
			assert.Equal(t, testcase.err, err)
		}
	}
}
