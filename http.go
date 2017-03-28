package thingfulx

import (
	"io/ioutil"
	"net/http"
)

// HTTPGetRequest is a simple helper that groups together common
// steps required to perform a GET requests.
// It returns a slice of bytes containing the actual HTTP response.
func HTTPGetRequest(urlString string, c Client) ([]byte, error) {

	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.DoHTTP(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if cerr := resp.Body.Close(); err == nil && cerr != nil {
			err = cerr
		}
	}()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, ErrNotFound
		}
		return nil, NewErrUnexpectedResponse(resp.Status)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}
