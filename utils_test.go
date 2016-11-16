package thingfulx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategorise(t *testing.T) {
	// testcases
	var testcases = []struct {
		inputs   []string
		expected Tag
	}{
		{
			[]string{
				"arduino air quality",
				"raspberry pi home sensor",
			},
			Tag("thingful:category=experiment"),
		},
		{
			[]string{
				"energy monitor",
				"home thermostat",
			},
			Tag("thingful:category=energy"),
		},
		{
			[]string{
				"healthkit",
				"exercise tracker",
			},
			Tag("thingful:category=health"),
		},
		{
			[]string{
				"garden moisture",
				"conservatory counter",
			},
			Tag("thingful:category=home"),
		},
		{
			[]string{
				"turtle neck",
				"shark week",
			},
			Tag("thingful:category=flora_and_fauna"),
		},
		{
			[]string{
				"weather checker",
				"flood monitor",
			},
			Tag("thingful:category=environment"),
		},
		{
			[]string{
				"truck",
				"plane or car",
			},
			Tag("thingful:category=transport"),
		},
		{
			[]string{
				"banana",
				"anything else",
			},
			Tag("thingful:category=miscellaneous"),
		},
	}

	var got Tag

	for _, testcase := range testcases {

		for _, input := range testcase.inputs {
			got = Categorise(input)

			assert.Equal(t, testcase.expected, got)
		}
	}
}
