package thingfulx

// Response is a struct encapsulating a response returned by a Fetcher
// instance. It contains the body of the response as a byte array, and the
// integer status code recording the status code of the HTTP response.
type Response struct {
	Body       []byte
	StatusCode int
}
