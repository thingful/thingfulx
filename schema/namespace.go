package schema

import (
	"strings"
)

const (
	// thingfulSchema is the base url for Thingful's ontology
	thingfulSchema = "http://purl.org/iot/vocab/thingful#"

	// m3liteSchema is the base url for the m3-lite ontology
	m3liteSchema = "http://purl.org/iot/vocab/m3-lite#"

	// ssnSchema is the base url for the SSN ontology
	ssnSchema = "http://purl.oclc.org/NET/ssnx/ssn#"

	// iotliteSchema is the base url for the IOT-LITE ontoloty
	iotliteSchema = "http://purl.oclc.org/NET/UNIS/fiware/iot-lite#"

	// xsdSchema is the base url for XSD terms
	xsdSchema = "http://www.w3.org/2001/XMLSchema#"

	// quSchema is the base url for QU ontologies
	quSchema = "http://purl.org/NET/ssnx/qu/qu#"

	// schemaOrg is base url for schema.org
	schemaOrgSchema = "http://schema.org/"

	// ccSchema is base url for Creative Commons REL
	ccSchema = "https://creativecommons.org/ns#"

	// jsonapiSchema is base url for JSONAPI schema
	jsonapiSchema = "http://purl.org/iot/vocab/jsonapi#"

	// thingfulQuSchema is base url for thingful quantity and unit schema
	thingfulQuSchema = "http://purl.org/iot/vocab/thingful/qu#"

	// biotopSchema is base url for Biotop schema, it is the domain for Biology and Medicine
	biotopSchema = "http://purl.org/biotop/biotop.owl#"

	// datexSchema is base url for Datex2 schema, it is the domain for Traffic & Transport
	datexSchema = "http://vocab.datex.org/terms#"

	// mobivocSchema is base url for Mobivoc schema, it is the domain for Mobility
	mobivocSchema = "http://schema.mobivoc.org/"

	// siriSchema is base url for unofficial SIRI (Service Interface for Real Time Information) ontology
	siriSchema = "http://purl.org/iot/vocab/siri#"

	// semSchema is the base url for the Simple Event Model ontology, used to
	// describe events in various subject domains.
	semSchema = "http://semanticweb.cs.vu.nl/2009/11/sem/"
)

var (
	// expanded is a map containing the schema to compact form mappings.
	namespaces = map[string]string{
		"thingful:":   thingfulSchema,
		"m3-lite:":    m3liteSchema,
		"ssn:":        ssnSchema,
		"iot-lite:":   iotliteSchema,
		"xsd:":        xsdSchema,
		"qu:":         quSchema,
		"schema:":     schemaOrgSchema,
		"cc:":         ccSchema,
		"jsonapi:":    jsonapiSchema,
		"thingfulqu:": thingfulQuSchema,
		"biotop:":     biotopSchema,
		"datex:":      datexSchema,
		"mobivoc:":    mobivocSchema,
		"siri:":       siriSchema,
		"sem:":        semSchema,
	}
)

// Expand returns the expanded version of a value (either prop or value). If we
// are unable to expand the value then we return the input string. This allows
// this function to be called safely on an already expanded property.
func Expand(val string) string {
	// k is the compact version, v is the expanded, and val is a compact string
	for k, v := range namespaces {
		if strings.HasPrefix(val, k) {
			return strings.Replace(val, k, v, 1)
		}
	}

	return val
}

// Compact is a function that converts an expanded form of a ontology property
// into its compact representation. If we are unable to compact the value then
// we return the input string. This allows this function to be called safely on
// an already compacted property.
func Compact(val string) string {
	// here k is still the compact, v is the expanded, but val is the expanded string
	for k, v := range namespaces {
		if strings.HasPrefix(val, v) {
			return strings.Replace(val, v, k, 1)
		}
	}

	return val
}
