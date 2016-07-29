package thingfulx

const (
	// ClientToken is the key we use when setting a token on the
	// Context Info property
	ClientToken = "client-token"
)

// Context is a map object information
// about the current environment that can be passed
// to a fetcher
type Context map[string]string

// Add is a helper method that stores a new value for a
// give key on the Context
func (c Context) Add(key, val string) {
	c[key] = val
}

// Del is a helper method that deletes a value associated
// with the key argument
func (c Context) Del(key string) {
	delete(c, key)
}
