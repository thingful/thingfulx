package thingfulx

import (
	"reflect"
	"testing"
)

func TestAddToContextInfo(t *testing.T) {
	c := Context{}
	c.Add("foo", "bar")
	c.Add("joe", "doe")

	expected := Context{"foo": "bar", "joe": "doe"}

	if !reflect.DeepEqual(c, expected) {
		t.Errorf("Unexpected error generating a new Context. Expectd %v, got %v", expected, c)
	}
}

func TestDeleteFromContextInfo(t *testing.T) {
	c := Context{}
	c.Add("foo", "bar")

	remove := "foo"

	c.Del(remove)

	if c[remove] != "" {
		t.Errorf("Unexpected error removing a Context property. '%s' value was not delete", remove)
	}
}
