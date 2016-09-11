package thingfulx

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCompleteThing(t *testing.T) {
	timestamp := time.Now()

	thing := Thing{
		Title:       "Title",
		Description: "Description",
		Category:    &Environment,
		Metadata: []Tag{
			Tag{
				Namespace: "urn:X-thingful",
				Predicate: "related",
				Value:     "citizen science,energy monitoring",
			},
			Tag{
				Namespace: "urn:X-thingful",
				Predicate: "provider",
				Value:     "energyco",
			},
		},
		Webpage:    "http://example.com/things",
		DataURL:    "http://example.com/things.json#id=123abc",
		IndexedAt:  timestamp,
		Lng:        -0.5,
		Lat:        55.5,
		Visibility: Open,
		RawData:    []byte("raw_data"),
		Channels: []Channel{
			Channel{
				Label:      "tempsensor1",
				Value:      "23.2",
				RecordedAt: timestamp,
				Metadata: []Tag{
					Tag{
						Namespace: "urn:X-thingful",
						Predicate: "unit",
						Value:     "celsius",
					},
				},
			},
			Channel{
				Label:      "energy",
				Value:      "467",
				RecordedAt: timestamp,
				Sensitive:  true,
				Metadata: []Tag{
					Tag{
						Namespace: "urn:X-thingful",
						Predicate: "unit",
						Value:     "kwh",
					},
				},
			},
		},
	}

	assert.Equal(t, thing.Title, "Title")
	assert.Equal(t, thing.Description, "Description")
	assert.Equal(t, thing.Category, &Environment)
	assert.Equal(t, thing.Webpage, "http://example.com/things")
	assert.Equal(t, thing.DataURL, "http://example.com/things.json#id=123abc")
	assert.Equal(t, thing.Lng, -0.5)
	assert.Equal(t, thing.Lat, 55.5)
	assert.Equal(t, thing.Visibility, Open)
	assert.Equal(t, thing.RawData, []byte("raw_data"))

	assert.Equal(t, len(thing.Metadata), 2)
	assert.Equal(t, len(thing.Channels), 2)
}
