package thingfulx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTag(t *testing.T) {
	testcases := []struct {
		input    string
		expected Tag
	}{
		{
			input: "temperature",
			expected: Tag{
				Namespace: "",
				Predicate: "",
				Value:     "temperature",
			},
		},
		{
			input: "thingful:quantity=temperature",
			expected: Tag{
				Namespace: "thingful",
				Predicate: "quantity",
				Value:     "temperature",
			},
		},
		{
			input: "geo:long=54.2",
			expected: Tag{
				Namespace: "geo",
				Predicate: "long",
				Value:     "54.2",
			},
		},
		{
			input: "*:long=54.2",
			expected: Tag{
				Namespace: "",
				Predicate: "",
				Value:     "*:long=54.2",
			},
		},
		{
			input: "*:*=54.2",
			expected: Tag{
				Namespace: "",
				Predicate: "",
				Value:     "*:*=54.2",
			},
		},
		{
			input: "thingful:title=\"foo bar\"",
			expected: Tag{
				Namespace: "thingful",
				Predicate: "title",
				Value:     "foo bar",
			},
		},
		{
			input: "*:long=",
			expected: Tag{
				Namespace: "",
				Predicate: "",
				Value:     "*:long=",
			},
		},
	}

	for _, testcase := range testcases {
		got := NewTag(testcase.input)
		assert.Equal(t, testcase.expected, got)
	}
}

func TestParseTag(t *testing.T) {
	testcases := []struct {
		input    string
		expected Tag
	}{
		{
			input: "temperature",
			expected: Tag{
				Namespace: "",
				Predicate: "",
				Value:     "temperature",
			},
		},
		{
			input: "thingful:quantity=temperature",
			expected: Tag{
				Namespace: "thingful",
				Predicate: "quantity",
				Value:     "temperature",
			},
		},
		{
			input: "geo:long=54.2",
			expected: Tag{
				Namespace: "geo",
				Predicate: "long",
				Value:     "54.2",
			},
		},
		{
			input: "*:long=54.2",
			expected: Tag{
				Namespace: "",
				Predicate: "long",
				Value:     "54.2",
			},
		},
		{
			input: "*:*=54.2",
			expected: Tag{
				Namespace: "",
				Predicate: "",
				Value:     "54.2",
			},
		},
		{
			input: "thingful:title=\"foo bar\"",
			expected: Tag{
				Namespace: "thingful",
				Predicate: "title",
				Value:     "foo bar",
			},
		},
		{
			input: "*:long=",
			expected: Tag{
				Namespace: "",
				Predicate: "long",
				Value:     "",
			},
		},
	}

	for _, testcase := range testcases {
		got := ParseTag(testcase.input)
		assert.Equal(t, testcase.expected, got)
	}
}

func TestNewMetadata(t *testing.T) {
	testcases := []struct {
		input    []string
		expected Metadata
	}{
		{
			input: []string{"temperature", "light"},
			expected: Metadata{
				{
					Namespace: "",
					Predicate: "",
					Value:     "temperature",
				},
				{
					Namespace: "",
					Predicate: "",
					Value:     "light",
				},
			},
		},
		{
			input: []string{"geo:long=54.2", "thingful:quantity=length", "foo"},
			expected: Metadata{
				{
					Namespace: "geo",
					Predicate: "long",
					Value:     "54.2",
				},
				{
					Namespace: "thingful",
					Predicate: "quantity",
					Value:     "length",
				},
				{
					Namespace: "",
					Predicate: "",
					Value:     "foo",
				},
			},
		},
	}

	for _, testcase := range testcases {
		got := NewMetadata(testcase.input...)
		assert.Equal(t, testcase.expected, got)
	}
}
