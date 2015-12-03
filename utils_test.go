package thingfulx

import (
	"reflect"
	"testing"
)

func TestCategorise(t *testing.T) {
	// testcases
	var testcases = []struct {
		inputs   []string
		expected *Category
	}{
		{
			[]string{
				"arduino air quality",
				"raspberry pi home sensor",
			},
			&Experiment,
		},
		{
			[]string{
				"energy monitor",
				"home thermostat",
			},
			&Energy,
		},
		{
			[]string{
				"healthkit",
				"exercise tracker",
			},
			&Health,
		},
		{
			[]string{
				"garden moisture",
				"conservatory counter",
			},
			&Home,
		},
		{
			[]string{
				"turtle neck",
				"shark week",
			},
			&Flora,
		},
		{
			[]string{
				"weather checker",
				"flood monitor",
			},
			&Environment,
		},
		{
			[]string{
				"truck",
				"plane or car",
			},
			&Transport,
		},
		{
			[]string{
				"banana",
				"anything else",
			},
			&Miscellaneous,
		},
	}

	var got *Category

	for _, testcase := range testcases {

		for _, input := range testcase.inputs {
			got = Categorise(input)

			if !reflect.DeepEqual(testcase.expected, got) {
				t.Errorf("Unexpected output, expected %#v, got %#v", testcase.expected, got)
			}
		}
	}
}
