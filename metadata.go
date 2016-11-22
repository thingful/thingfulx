package thingfulx

import (
	"regexp"
)

const (
	// WildCard is our machinetag wild card character.
	WildCard = "*"
)

var (
	// machineTagPattern is a compiled regexp we use to identify a machinetag
	// input vs a regular tag
	machineTagPattern = regexp.MustCompile("^([a-z]\\w+|\\*)(:(([a-z]\\w+)|\\*))=(([^\\s]+)|\"([^\"]+)\")?$")
)

// Tag represents a metadata fact we can state about either a Thing or a
// channel.
type Tag struct {
	Namespace string
	Predicate string
	Value     string
}

// Metadata is a type alias for a slice of Tags
type Metadata []Tag

// NewTag is a constructor function for making a new Tag instance from a given
// string. Some example tags might be:
//
// * temperature
// * light
//
// However really these are intended for structured tags that look like this:
//
// * thingful:quantity=temperature
// * thingful:unit=lux
func NewTag(rawTag string) Tag {
	var namespace, predicate, value string

	if machineTagPattern.MatchString(rawTag) {
		matches := machineTagPattern.FindStringSubmatch(rawTag)

		namespace = convertWildcard(matches[1])
		predicate = convertWildcard(matches[3])
		if matches[6] != "" {
			value = matches[6]
		} else {
			value = matches[7]
		}
	} else {
		value = rawTag
	}

	return Tag{
		Namespace: namespace,
		Predicate: predicate,
		Value:     value,
	}
}

// convertWildcard converts a machinetag wildcard character into an empty
// string.
func convertWildcard(val string) string {
	if val == WildCard {
		return ""
	}
	return val
}

// NewMetadata takes one or more rawTags as variadic string parameters, and
// returns an instantiated Metadata instance.
func NewMetadata(rawTags ...string) Metadata {
	metadata := Metadata{}

	for _, rawTag := range rawTags {
		metadata = append(metadata, NewTag(rawTag))
	}

	return metadata
}
