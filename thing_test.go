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
		Webpage:     "http://example.com/things",
		DataURL:     "http://example.com/things.json#id=123abc",
		IndexedAt:   timestamp,
		Lng:         -0.5,
		Lat:         55.5,
		Visibility:  Open,
		RawData:     []byte("raw_data"),
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
}
