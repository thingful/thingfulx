package thingfulx

import (
	"regexp"
)

const (
	// WildCard is our machinetag wild card character.
	WildCard = "*"
)

var (
	// newTagPattern is a "strict" pattern that will only match a complete
	// machinetag.
	newTagPattern = regexp.MustCompile("^([a-z]\\w+)(:(([a-z]\\w+)))=(([^\\s]+)|\"([^\"]+)\")?$")

	// parseTagPattern is a more lenient pattern intended for parsing partial
	// machinetags used when searching.
	parseTagPattern = regexp.MustCompile("^([a-z]\\w+|\\*)(:(([a-z]\\w+)|\\*))=(([^\\s]+)|\"([^\"]+)\")?$")
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

// NewTag creates a new tag that will either be a full complete machinetag with
// namespace and predicate non-empty, or all data will be captured as the
// value.
func NewTag(rawTag string) Tag {
	var namespace, predicate, value string

	if newTagPattern.MatchString(rawTag) {
		matches := newTagPattern.FindStringSubmatch(rawTag)

		namespace = matches[1]
		predicate = matches[3]
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

// ParseTag leniently constructs a Tag representation, with wildcard values for
// namespace or predicate if not present.
func ParseTag(rawTag string) Tag {
	var namespace, predicate, value string

	if parseTagPattern.MatchString(rawTag) {
		matches := parseTagPattern.FindStringSubmatch(rawTag)

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
