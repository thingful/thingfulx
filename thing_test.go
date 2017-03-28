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
		Webpage:     "http://example.com/things",
		Endpoint: &Endpoint{
			URL:            "http://example.com/things.json#id=123abc",
			ContentType:    "application/exe",
			Authentication: "thingful:Open",
		},
		IndexedAt: timestamp,
		Location: &Location{
			Lng:       -0.5,
			Lat:       55.5,
			Elevation: 123.0,
		},
		Provider: &Provider{
			ID:          "testprovider",
			Name:        "Test Provider",
			Description: "Test Provider description",
			URL:         "http://example.com",
		},
		Metadata: []Metadata{
			{
				Prop: "thingful:Category",
				Val:  Environment.Name,
			},
			{
				Prop: "thingful:CategoryUID",
				Val:  Environment.UID,
			},
		},
		Channels: []Channel{
			{
				ID: "value1",
				Metadata: []Metadata{
					{
						Prop: "baz",
						Val:  "qux",
					},
				},
				Observations: []Observation{
					{
						RecordedAt: timestamp,
						Location: &Location{
							Lng: -1.23,
							Lat: 45.6,
						},
						Data: &Data{
							Type: "xsd:double",
							Val:  43.3,
						},
					},
				},
			},
		},
	}

	assert.Equal(t, thing.Title, "Title")
	assert.Equal(t, thing.Description, "Description")
	assert.Equal(t, thing.Webpage, "http://example.com/things")
	assert.Equal(t, thing.Endpoint.URL, "http://example.com/things.json#id=123abc")
	assert.Equal(t, thing.Endpoint.ContentType, "application/exe")
	assert.Equal(t, thing.Endpoint.Authentication, "thingful:Open")
	assert.Equal(t, thing.Location.Lng, -0.5)
	assert.Equal(t, thing.Location.Lat, 55.5)
	assert.Equal(t, thing.Location.Elevation, 123.0)
	assert.Equal(t, thing.Provider.ID, "testprovider")
	assert.Equal(t, thing.Provider.Name, "Test Provider")
	assert.Equal(t, thing.Provider.Description, "Test Provider description")
	assert.Equal(t, thing.Provider.URL, "http://example.com")
	assert.Equal(t, thing.Metadata[0].Prop, "thingful:Category")
	assert.Equal(t, thing.Metadata[0].Val, "Environment")
	assert.Equal(t, thing.Channels[0].ID, "value1")
	assert.Equal(t, thing.Channels[0].Observations[0].RecordedAt, timestamp)
	assert.Equal(t, thing.Channels[0].Observations[0].Location.Lat, 45.6)
	assert.Equal(t, thing.Channels[0].Observations[0].Location.Lng, -1.23)
	assert.Equal(t, thing.Channels[0].Observations[0].Data.Type, "xsd:double")
	assert.Equal(t, thing.Channels[0].Observations[0].Data.Val, 43.3)
}

func TestGetMetaValue(t *testing.T) {
	thing := Thing{
		Title:       "Title",
		Description: "Description",
		Webpage:     "http://example.com/things",
		Endpoint: &Endpoint{
			URL:            "http://example.com/things.json#id=123abc",
			ContentType:    "application/exe",
			Authentication: "thingful:Open",
		},
		Metadata: []Metadata{
			{
				Prop: "thingful:Category",
				Val:  Environment.Name,
			},
			{
				Prop: "thingful:hasRoutes",
				Val:  "1,2,3",
			},
		},
		Channels: []Channel{
			{
				ID: "value1",
				Metadata: []Metadata{
					{
						Prop: "baz",
						Val:  "qux,boo",
					},
				},
			},
		},
	}

	testcases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "thingful:Category",
			expected: []string{"Environment"},
		},
		{
			input:    "thingful:hasRoutes",
			expected: []string{"1", "2", "3"},
		},
		{
			input:    "xxx",
			expected: nil,
		},
	}

	for _, testcase := range testcases {
		got := thing.MetaValForProp(testcase.input)
		assert.Equal(t, testcase.expected, got)
	}
}

func TestGetChannelsMetaValue(t *testing.T) {
	thing := Thing{
		Title:       "Title",
		Description: "Description",
		Webpage:     "http://example.com/things",
		Endpoint: &Endpoint{
			URL:            "http://example.com/things.json#id=123abc",
			ContentType:    "application/exe",
			Authentication: "thingful:Open",
		},
		Metadata: []Metadata{
			{
				Prop: "thingful:Category",
				Val:  Environment.Name,
			},
			{
				Prop: "thingful:CategoryUID",
				Val:  Environment.UID,
			},
		},
		Channels: []Channel{
			{
				ID: "value1",
				Metadata: []Metadata{
					{
						Prop: "baz",
						Val:  "qux,boo",
					},
					{
						Prop: "meta",
						Val:  "test",
					},
				},
			},
		},
	}

	testcases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "baz",
			expected: []string{"qux", "boo"},
		},
		{
			input:    "meta",
			expected: []string{"test"},
		},
		{
			input:    "xxx",
			expected: nil,
		},
	}

	for _, testcase := range testcases {
		got := thing.ChannelMetaValForProp(testcase.input)
		assert.Equal(t, testcase.expected, got)
	}
}
