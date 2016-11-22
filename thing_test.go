package thingfulx

import (
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCompleteThing(t *testing.T) {
	timestamp := time.Now()

	webpage, _ := url.Parse("http://example.com/things")
	dataURL, _ := url.Parse("http://api.example.com/things/123abc.json")

	thing := Thing{
		Title:       "Title",
		Description: "Description",
		Webpage:     webpage,
		DataURL:     dataURL,
		IndexedAt:   timestamp,
		Location: &Location{
			Longitude: -0.5,
			Latitude:  55.5,
		},
		Metadata: NewMetadata(
			"thingful:category=environment",
			"thingful:visibility=open",
		),
		Category:   &Environment,
		Visibility: Open,
		Channels: []Channel{
			Channel{
				Value:      "17.2",
				Type:       NumberType,
				RecordedAt: timestamp,
				Metadata: NewMetadata(
					"thingful:property=temperature",
					"thingful:unit=celsius",
					"thingful:measurementType=interval",
					"thingful:label=temp_sensor1",
				),
			},
			Channel{
				Value:      "cloudy",
				Type:       StringType,
				RecordedAt: timestamp,
				Metadata: NewMetadata(
					"thingful:measurementType=nominal",
					"thingful:label=weather_desc",
				),
			},
		},
	}

	assert.Equal(t, "Title", thing.Title)
	assert.Equal(t, "Description", thing.Description)
	assert.Equal(t, "http://example.com/things", thing.Webpage.String())
	assert.Equal(t, "http://api.example.com/things/123abc.json", thing.DataURL.String())
	assert.Equal(t, Longitude(-0.5), thing.Location.Longitude)
	assert.Equal(t, Latitude(55.5), thing.Location.Latitude)
	assert.Len(t, thing.Metadata, 2)
	assert.Len(t, thing.Channels, 2)
	channel := thing.Channels[0]
	assert.Equal(t, "17.2", channel.Value)
	assert.Equal(t, NumberType, channel.Type)
	assert.Equal(t, timestamp, channel.RecordedAt)
	assert.Len(t, channel.Metadata, 4)
	assert.Contains(t, channel.Metadata, Tag{Namespace: "thingful", Predicate: "unit", Value: "celsius"})
	assert.Equal(t, &Environment, thing.Category)
	assert.Equal(t, Open, thing.Visibility)
}
