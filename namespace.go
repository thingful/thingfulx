package thingfulx

import "strings"

const (
	thingfulSchema = Ontology("https://thingful.github.io/schema#")

	m3liteSchema = Ontology("http://purl.org/iot/vocab/m3-lite#")

	ssnSchema = Ontology("http://purl.oclc.org/NET/ssnx/ssn#")

	iotliteSchema = Ontology("http://purl.oclc.org/NET/UNIS/fiware/iot-lite#")

	xsdSchema = Ontology("http://www.w3.org/2001/XMLSchema#")
)

// Ontology type is used to define a set list of ontologies known
// by the Thingfulx library.
type Ontology string

// String is the Ontology type stringer implementation
func (o Ontology) String() string {
	return string(o)
}

var namespace = map[string]Ontology{
	"thingful:": thingfulSchema,
	"m3-lite:":  m3liteSchema,
	"ssn:":      ssnSchema,
	"iot-lite:": iotliteSchema,
	"xsd:":      xsdSchema,
}

// ExpandNamespace reruns the expanded version of n.
func ExpandNamespace(n string) Ontology {
	for k, v := range namespace {
		if strings.Contains(n, k) {
			j := strings.Replace(n, k, string(v), 1)
			return Ontology(j)
		}
	}

	return ""
}
