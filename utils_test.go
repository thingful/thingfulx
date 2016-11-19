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
			NewTag("thingful:category=experiment"),
		},
		{
			[]string{
				"energy monitor",
				"home thermostat",
			},
			NewTag("thingful:category=energy"),
		},
		{
			[]string{
				"healthkit",
				"exercise tracker",
			},
			NewTag("thingful:category=health"),
		},
		{
			[]string{
				"garden moisture",
				"conservatory counter",
			},
			NewTag("thingful:category=home"),
		},
		{
			[]string{
				"turtle neck",
				"shark week",
			},
			NewTag("thingful:category=flora_and_fauna"),
		},
		{
			[]string{
				"weather checker",
				"flood monitor",
			},
			NewTag("thingful:category=environment"),
		},
		{
			[]string{
				"truck",
				"plane or car",
			},
			NewTag("thingful:category=transport"),
		},
		{
			[]string{
				"banana",
				"anything else",
			},
			NewTag("thingful:category=miscellaneous"),
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
