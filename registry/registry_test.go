package registry_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/thingful/thingfulx"
	"github.com/thingful/thingfulx/registry"
)

type testIndexer struct{}

func (testIndexer) Provider() *thingfulx.Provider { return &thingfulx.Provider{Uid: "testindexer"} }
func (testIndexer) URLS(ctx context.Context, client thingfulx.Client, delay time.Duration, out chan<- thingfulx.URLData) {
}
func (testIndexer) Fetch(ctx context.Context, urlStr string, client thingfulx.Client, clock thingfulx.Clock) ([]byte, error) {
	return nil, errors.New("No data available")
}
func (testIndexer) Parse(rawData []byte, urlStr string, clock thingfulx.Clock) ([]*thingfulx.Channel, error) {
	return nil, errors.New("Unable to parse")
}

func newTestIndexer() (thingfulx.Indexer, error) {
	return &testIndexer{}, nil
}

func TestRegistry(t *testing.T) {
	r := registry.NewRegistry()
	err := r.Register(newTestIndexer)
	assert.Nil(t, err)

	indexer, err := r.GetIndexer("testindexer")
	assert.Nil(t, err)
	assert.NotNil(t, indexer)
	assert.Equal(t, "testindexer", indexer.Provider().GetUid())

	_, err = r.GetIndexer("unknownprovider")
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, registry.ErrUnknownProvider))

	indexers := r.GetIndexers()
	assert.Len(t, indexers, 1)
}
