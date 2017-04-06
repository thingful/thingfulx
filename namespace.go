package thingfulx

import (
	"errors"
	"strings"
)

const (
	// thingfulSchema is the base url for Thingful's ontology
	thingfulSchema = "https://thingful.github.io/schema#"

	// m3liteSchema is the base url for the m3-lite ontology
	m3liteSchema = "http://purl.org/iot/vocab/m3-lite#"

	// ssnSchema is the base url for the SSN ontology
	ssnSchema = "http://purl.oclc.org/NET/ssnx/ssn#"

	// iotliteSchema is the base url for the IOT-LITE ontoloty
	iotliteSchema = "http://purl.oclc.org/NET/UNIS/fiware/iot-lite#"

	// xsdSchema is the base url for XSD terms
	xsdSchema = "http://www.w3.org/2001/XMLSchema#"
)

var (
	// ErrUnknownNamespace is an error returned when we attempt to expand a
	// property that we don't have a namespace mapping for
	ErrUnknownNamespace = errors.New("thingfulx: Unknown namespace")
)

var (
	// expanded is a map containing the schema to compact form mappings.
	namespaces = map[string]string{
		"thingful:": thingfulSchema,
		"m3-lite:":  m3liteSchema,
		"ssn:":      ssnSchema,
		"iot-lite:": iotliteSchema,
		"xsd:":      xsdSchema,
	}
)

// Expand returns the expanded version of a value (either prop or
// value).
func Expand(val string) (string, error) {
	// k is the compact version, v is the expanded, and val is a compact string
	for k, v := range namespaces {
		if strings.Contains(val, k) {
			return strings.Replace(val, k, v, 1), nil
		}
	}

	return "", ErrUnknownNamespace
}

// Compact is a function that converts an expanded form of a ontology property
// into its compact representation.  This function returns an error if we are
// unable to match the namespace.
func Compact(val string) (string, error) {
	// here k is still the compact, v is the expanded, but val is the expanded string
	for k, v := range namespaces {
		if strings.Contains(val, v) {
			return strings.Replace(val, v, k, 1), nil
		}
	}

	return "", ErrUnknownNamespace
}
