package thingfulx

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/thingful/thingfulx/schema"
)

func TestCompleteThing(t *testing.T) {
	timestamp := time.Now()

	thing := Thing{
		Title:       "Title",
		Description: "Description",
		Webpage:     "http://example.com/things",
		Visibility:  Open,
		Endpoint: &Endpoint{
			URL:         "http://example.com/things.json#id=123abc",
			ContentType: "application/exe",
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
		ThingType:       schema.Expand("thingful:ConnectedDevice"),
		DataLicense:     GetDataLicense(CCByV3URL),
		AttributionName: "Test Data Author",
		AttributionURL:  "http://some/data/author/url",
		Category:        Environment.Name,
		UpdateInterval:  900,
		Metadata: []Metadata{
			{
				Prop: schema.HasCategory,
				Val:  Environment.Name,
			},
			{
				Prop: schema.HasCategoryUID,
				Val:  Environment.UID,
			},
		},
		Channels: []Channel{
			{
				ID:           "value1",
				QuantityKind: schema.Expand("m3-lite:WindSpeed"),
				Unit:         schema.Expand("m3-lite:KilometerPerHour"),
				MeasuredBy:   schema.Expand("m3-lite:WindSpeedSensor"),
				DomainOfInterest: []string{
					schema.Expand("m3-lite:Weather"),
					schema.Expand("m3-lite:Environment"),
				},
				Metadata: []Metadata{
					{
						Prop: "baz",
						Val:  "qux",
					},
				},
				Type: schema.DoubleType,
				Observations: []Observation{
					{
						RecordedAt: timestamp,
						Location: &Location{
							Lng: -1.23,
							Lat: 45.6,
						},
						Val: 43.3,
					},
				},
			},
		},
	}

	assert.Equal(t, thing.Title, "Title")
	assert.Equal(t, thing.Description, "Description")
	assert.Equal(t, thing.Webpage, "http://example.com/things")
	assert.Equal(t, thing.Visibility, Visibility("open"))
	assert.Equal(t, thing.Endpoint.URL, "http://example.com/things.json#id=123abc")
	assert.Equal(t, thing.Endpoint.ContentType, "application/exe")
	assert.Equal(t, thing.Location.Lng, -0.5)
	assert.Equal(t, thing.Location.Lat, 55.5)
	assert.Equal(t, thing.Location.Elevation, 123.0)
	assert.Equal(t, thing.Provider.ID, "testprovider")
	assert.Equal(t, thing.Provider.Name, "Test Provider")
	assert.Equal(t, thing.Provider.Description, "Test Provider description")
	assert.Equal(t, thing.Provider.URL, "http://example.com")
	assert.Equal(t, thing.ThingType, schema.Expand("thingful:ConnectedDevice"))
	assert.Equal(t, thing.DataLicense, GetDataLicense(CCByV3URL))
	assert.Equal(t, thing.AttributionName, "Test Data Author")
	assert.Equal(t, thing.AttributionURL, "http://some/data/author/url")
	assert.Equal(t, thing.Category, Environment.Name)
	assert.Equal(t, thing.UpdateInterval, 900)
	assert.Equal(t, thing.Metadata[0].Prop, schema.HasCategory)
	assert.Equal(t, thing.Metadata[0].Val, "Environment")
	assert.Len(t, thing.Channels, 1)
	assert.Equal(t, thing.Channels[0].ID, "value1")
	assert.Equal(t, thing.Channels[0].Type, schema.DoubleType)
	assert.Equal(t, thing.Channels[0].QuantityKind, schema.Expand("m3-lite:WindSpeed"))
	assert.Equal(t, thing.Channels[0].Unit, schema.Expand("m3-lite:KilometerPerHour"))
	assert.Equal(t, thing.Channels[0].MeasuredBy, schema.Expand("m3-lite:WindSpeedSensor"))
	assert.Equal(t, thing.Channels[0].DomainOfInterest[0], schema.Expand("m3-lite:Weather"))
	assert.Equal(t, thing.Channels[0].Observations[0].RecordedAt, timestamp)
	assert.Equal(t, thing.Channels[0].Observations[0].Location.Lat, 45.6)
	assert.Equal(t, thing.Channels[0].Observations[0].Location.Lng, -1.23)
	assert.Equal(t, thing.Channels[0].Observations[0].Val, 43.3)
}

func TestMetaVal(t *testing.T) {
	thing := Thing{
		Title:       "Title",
		Description: "Description",
		Webpage:     "http://example.com/things",
		Visibility:  Open,
		Endpoint: &Endpoint{
			URL:         "http://example.com/things.json#id=123abc",
			ContentType: "application/json",
		},
		Metadata: []Metadata{
			{
				Prop: schema.HasCategory,
				Val:  Environment.Name,
			},
			{
				Prop: schema.Expand("thingful:hasRoute"),
				Val:  "1",
			},
			{
				Prop: schema.Expand("thingful:hasRoute"),
				Val:  "2",
			},
			{
				Prop: schema.Expand("thingful:hasRoute"),
				Val:  "3",
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
			},
		},
	}

	testcases := []struct {
		input    string
		expected string
	}{
		{
			input:    schema.HasCategory,
			expected: "Environment",
		},
		{
			input:    schema.Expand("thingful:hasRoute"),
			expected: "1",
		},
		{
			input:    "xxx",
			expected: "",
		},
	}

	for _, testcase := range testcases {
		got := thing.MetaVal(testcase.input)
		assert.Equal(t, testcase.expected, got)
	}
}
