package registry

import (
	"errors"
	"fmt"
	"sync"

	"github.com/thingful/thingfulx"
)

const (
	// ErrUnknownProvider is a sentinel error type returned if a client requests
	// an unknown indexer from the registry.
	ErrUnknownProvider = thingfulx.Error("thingfulx: unknown data provider")
)

// Registry is a threadsafe data structure for holding instantiated
// thingfulx.Indexers, and returning them to a caller.
type Registry struct {
	sync.RWMutex
	indexers map[string]thingfulx.Indexer
}

// NewRegistry is a constructor function that returns a Registry ready for use.
func NewRegistry() *Registry {
	return &Registry{
		indexers: map[string]thingfulx.Indexer{},
	}
}

// Register is a function that takes as input a function that is capable of
// contructing a new indexer instance. It then instantiates the indexer
// returning any error, but should the operation succeed the indexer is saved
// into a registry for later use.
func (r *Registry) Register(builder func() (thingfulx.Indexer, error)) error {
	indexer, err := builder()
	if err != nil {
		return fmt.Errorf("failed to instantiate indexer: %v", err)
	}

	if indexer.Provider() == nil {
		return errors.New("indexer has failed to provide data provider information")
	}

	r.Lock()
	defer r.Unlock()

	r.indexers[indexer.Provider().GetUid()] = indexer

	return nil
}

// MustRegister wraps our Register function but panics on any error as we cannot
// proceed.
func (r *Registry) MustRegister(builder func() (thingfulx.Indexer, error)) {
	err := r.Register(builder)
	if err != nil {
		panic(err)
	}
}

// GetIndexer attempts to return the indexer for the specified provider UID.
func (r *Registry) GetIndexer(providerUID string) (thingfulx.Indexer, error) {
	r.RLock()
	defer r.RUnlock()

	if indexer, ok := r.indexers[providerUID]; ok {
		return indexer, nil
	}

	return nil, fmt.Errorf("Unknown indexer '%s': %w", providerUID, ErrUnknownProvider)
}

// GetIndexers returns a list of the registered Indexers in the registry
func (r *Registry) GetIndexers() []thingfulx.Indexer {
	r.RLock()
	defer r.RUnlock()

	indexers := []thingfulx.Indexer{}
	for _, i := range r.indexers {
		indexers = append(indexers, i)
	}

	return indexers
}
