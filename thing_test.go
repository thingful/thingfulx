package thingfulx

import (
	"testing"
	"time"
)

func TestThingStringer(t *testing.T) {
	expected := "Thing<Title:Weather,URI:http://example.com/pages/foo>"

	thing := Thing{
		Title: "Weather",
		URI:   "http://example.com/pages/foo",
	}

	if thing.String() != expected {
		t.Errorf("Unexpected String() output, expected '%s', got '%s'", expected, thing.String())
	}
}

func TestCompleteThing(t *testing.T) {
	timestamp := time.Now()

	thing := Thing{
		Title:       "Title",
		Description: "Description",
		Category:    Environment,
		Webpage:     "http://example.com/things",
		URI:         "http://example.com/things#123abc",
		IndexedAt:   timestamp,
		Lng:         -0.5,
		Lat:         55.5,
		Visibility:  Public,
		DataURL:     "http://example.com/things.json",
	}

	if thing.Title != "Title" {
		t.Errorf("Unexpected value, expected 'Title', got '%s'", thing.Title)
	}

	if thing.Description != "Description" {
		t.Errorf("Unexpected value, expected 'Description', got '%s'", thing.Description)
	}

	if thing.Category != Environment {
		t.Errorf("Unexpected value, expected 'environment', got '%s'", thing.Category)
	}

	if thing.Webpage != "http://example.com/things" {
		t.Errorf("Unexpected value, expected 'http://example.com/things', got '%s'", thing.Webpage)
	}

	if thing.URI != "http://example.com/things#123abc" {
		t.Errorf("Unexpected value, expected 'http://example.com/things#123abc', got '%s'", thing.URI)
	}

	if thing.Lng != -0.5 {
		t.Errorf("Unexpected value, expected '-0.5', got '%s'", thing.Lng)
	}

	if thing.Lat != 55.5 {
		t.Errorf("Unexpected value, expected '55.5', got '%s'", thing.Lat)
	}

	if thing.Visibility != Public {
		t.Errorf("Unexpected value, expected 'public', got '%s'", thing.Visibility)
	}

	if thing.DataURL != "http://example.com/things.json" {
		t.Errorf("Unexpected value, expected 'http://example.com/things.json', got '%s'", thing.DataURL)
	}
}
