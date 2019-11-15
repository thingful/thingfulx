package thingfulx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sync"

	"github.com/temoto/robotstxt"
)

// permissiveString contains a robots.txt file that allows all access. Used when
// we cannot obtain a valid robots.txt from a domain.
const permissiveString = "User-Agent: *\nDisallow:"

var (
	// robotsPath is a URL containing the well known path for robots.txt files
	robotsPath, _ = url.Parse("/robots.txt")
)

func newRobots() *robots {
	return &robots{
		cache: map[string]*robotstxt.Group{},
	}
}

// robots is a type that provides a method that checks whether a request is
// permitted to be sent to a domain - either because their robots.txt explicitly
// permits access, or that we are unable to obtain a valid robots.txt file from
// the well known path. robots type is thread safe and expected to be able to be
// called from multiple goroutines. Maintains a simple in memory cache of
// received robots.txt responses to avoid making repeated calls.
type robots struct {
	sync.RWMutex
	cache map[string]*robotstxt.Group
}

// permitted is a function on our robots instance that allows the caller to
// obtain a robots.txt agent and use that agent to test whether the client
// should be permitted to access the resource URL. Internally the robots
// instance maintains an in memory cache of parsers meaning that once running it
// will not keep requesting the remote robots.txt content for a specific host.
// If no robots.txt is available we use a permissive check that allows all
// requests.
func (r *robots) permitted(ctx context.Context, resourceURL *url.URL, userAgent string, client *http.Client) bool {
	group := r.getGroup(ctx, resourceURL, userAgent, client)
	return group.Test(resourceURL.Path)
}

// getGroup attempts to obtain a *robotstxt.Group instance for the specified
// URL. We cache any response here to speed subsequent requests. If any error
// happens on retrieiving a robots.txt, we assume permitted.
func (r *robots) getGroup(ctx context.Context, resourceURL *url.URL, userAgent string, client *http.Client) *robotstxt.Group {
	// look in cache and return if present
	group := r.getGroupFromCache(resourceURL.Host)
	if group != nil {
		return group
	}

	// if not present then attempt to fetch from remote
	group, err := r.fetchGroup(ctx, resourceURL, userAgent, client)
	if err != nil {
		// on any error return our fully permissive group
		group = permissiveGroup(userAgent)
	}

	// update cache
	r.Lock()
	defer r.Unlock()
	r.cache[resourceURL.Host] = group

	return group
}

// getGroupFromCache attempts to return our robotstxt.Group instance from the in
// memory cache (a map). Returns nil if an instance for the specified host is
// not present
func (r *robots) getGroupFromCache(host string) *robotstxt.Group {
	r.RLock()
	defer r.RUnlock()

	group, ok := r.cache[host]
	if ok {
		return group
	}

	return nil
}

// fetchGroup attempts to fetch and parse the robots.txt content from the remote
// domain. On any error we return a fully permissive robotstxt.Group meaning
// requests will be permitted until we check again.
func (r *robots) fetchGroup(ctx context.Context, resourceURL *url.URL, userAgent string, client *http.Client) (*robotstxt.Group, error) {
	robotsURL := resourceURL.ResolveReference(robotsPath)

	req, err := http.NewRequest(http.MethodGet, robotsURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to build request for robots.txt: %v", err)
	}

	// propagate our context onwards
	req = req.WithContext(ctx)

	req.Header.Add("User-Agent", userAgent)

	// make request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to make request for robots.txt: %v", err)
	}

	// ensure response body closed as robotstxt doesn't do this for us
	defer resp.Body.Close()

	// attempt to parse robotstxt Group from the response
	robotsData, err := robotstxt.FromResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse robots.txt response: %v", err)
	}

	// Attempt to obtain robotstxt.Group from the parsed response
	group := robotsData.FindGroup(userAgent)
	if group == nil {
		return nil, errors.New("Failed to find matching group")
	}

	return group, nil
}

// permissiveGroup is a helper function that returns a fully permissive
// robotstxt.Group which we apply if we are unable to obtain an actual
// robots.txt from the remote data provider.
func permissiveGroup(userAgent string) *robotstxt.Group {
	rd, _ := robotstxt.FromString(permissiveString)
	return rd.FindGroup(userAgent)
}
