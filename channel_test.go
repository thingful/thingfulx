package thingfulx

import (
	"testing"
	"time"
)

func TestChannelStringer(t *testing.T) {
	c := Channel{
		ID:         "temp",
		Value:      "23.2",
		RecordedAt: time.Now(),
		Unit: Unit{
			Name:   "gram",
			Symbol: "g",
		},
	}

	if c.String() != "Channel<ID:temp,Value:23.2>" {
		t.Errorf("Unexpected String() output, expected 'Channel<ID:temp,Value:23.2>', got '%s'", c.String())
	}
}
