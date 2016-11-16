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
		Metadata: []Tag{
			"thingful:category=environment",
			"thingful:visibility=open",
		},
		RawData: []byte("raw_data"),
		Channels: []Channel{
			Channel{
				Value:      "17.2",
				Type:       NumberType,
				RecordedAt: timestamp,
				Metadata: []Tag{
					"thingful:property=temperature",
					"thingful:unit=celsius",
					"thingful:measurementType=interval",
					"thingful:label=temp_sensor1",
				},
			},
			Channel{
				Value:      "cloudy",
				Type:       StringType,
				RecordedAt: timestamp,
				Metadata: []Tag{
					"thingful:measurementType=nominal",
					"thingful:label=weather_desc",
				},
			},
		},
	}

	assert.Equal(t, "Title", thing.Title)
	assert.Equal(t, "Description", thing.Description)
	assert.Equal(t, "http://example.com/things", thing.Webpage.String())
	assert.Equal(t, "http://api.example.com/things/123abc.json", thing.DataURL.String())
	assert.Equal(t, -0.5, thing.Location.Longitude)
	assert.Equal(t, 55.5, thing.Location.Latitude)
	assert.Equal(t, thing.RawData, []byte("raw_data"))
	assert.Contains(t, thing.Metadata, Tag("thingful:category=environment"))
	assert.Contains(t, thing.Metadata, Tag("thingful:visibility=open"))
	assert.Len(t, thing.Channels, 2)

	channel := thing.Channels[0]
	assert.Equal(t, "17.2", channel.Value)
	assert.Equal(t, NumberType, channel.Type)
	assert.Equal(t, timestamp, channel.RecordedAt)
	assert.Contains(t, channel.Metadata, Tag("thingful:property=temperature"))
}
