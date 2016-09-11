package thingfulx

// Tag represents a metadata type fact that can be recorded about a thing or a
// channel. A single Tag represents a single fact about the described resource
// and comprises an optional Namespace plus a mandatory Predicate and Value
// pair. An example Tag might be:
//
//	tag := &Tag{
//		Namespace: "urn:X-thingful:rels",
//		Predicate: "unit",
//		Value:     "celsius",
//	}
type Tag struct {
	Namespace string
	Predicate string
	Value     string
}
