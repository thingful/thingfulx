package thingfulx

import (
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/temoto/robotstxt"
)

const (
	// ClientToken is the key we use when setting a token on the
	// client clientInfo property
	ClientToken = "client-token"

	// permissiveString is a permissive robots.txt specification to be used when
	// none is available from the upstream source. It allows access to all
	// resources.
	permissiveString = "User-agent: *\nDisallow:"
)

var (
	// robotsTxtPath is a url.URL parsed path for robots.txt file. This cannot return an error
	robotsTxtPath, _ = url.Parse("/robots.txt")
)

// Client specifies an interface we will use for making outgoing requests
// within each Fetcher. Currently everything is HTTP, but that won't
// necessarily remain the case, so we add a layer of abstraction meaning we can
// add methods to our Client interface to support other protocols.
type Client interface {
	// DoHTTP is a thin wrapper around the `Do` method on `net/http.Client`. It
	// simply ensures we set the correct user agent header before making the
	// request
	DoHTTP(req *http.Request) (*http.Response, error)
}

// NewClient is a constructor function that instantiates and returns a new
// client struct, ensuring it sets `timeout` properly on the returned
// http.Client. This constructor also allows passing a list of explicitly
// permitted hosts via a variadic final argument. This is for those cases where
// we have explicit permission to index a data provider despite their
// robots.txt forbidding access.
func NewClient(userAgent string, timeout time.Duration, permittedHosts ...string) Client {
	robotsData, _ := robotstxt.FromString(permissiveString)

	permittedHostsMap := map[string]bool{}

	for _, host := range permittedHosts {
		permittedHostsMap[host] = true
	}

	return &client{
		userAgent: userAgent,
		httpClient: &http.Client{
			Timeout: timeout,
		},
		permittedHosts:  permittedHostsMap,
		robotAgents:     map[string]*robotstxt.Group{},
		permissiveAgent: robotsData.FindGroup(userAgent),
	}
}

// client is our concrete implementation of the Client interface. As well as
// the http client and the user agent string, this struct also contains some
// machinery used when checking remote sites robots.txt files to see whether we
// are allowed to access.
type client struct {
	sync.RWMutex
	userAgent       string
	httpClient      *http.Client
	permittedHosts  map[string]bool
	robotAgents     map[string]*robotstxt.Group
	permissiveAgent *robotstxt.Group
}

// DoHTTP is a thin wrapper around the `Do` method on `net/http.Client`. It
// simply ensures we set the correct user agent header before making the
// request
func (c *client) DoHTTP(req *http.Request) (*http.Response, error) {
	if c.permittedToFetch(req.URL) {
		// set the user agent here
		req.Header.Add("User-Agent", c.userAgent)

		return c.httpClient.Do(req)
	}

	return nil, ErrNotPermittedToFetch
}

func (c *client) permittedToFetch(u *url.URL) bool {
	agent := c.getAgent(u)

	return c.permittedHosts[u.Host] || agent.Test(u.Path)
}

func (c *client) getAgent(u *url.URL) *robotstxt.Group {
	c.Lock()
	defer c.Unlock()

	robotAgent, found := c.robotAgents[u.Host]

	if found {
		return robotAgent
	}

	robotAgent, err := c.createRobotAgent(u)

	// if we get an error for a robots.txt response we return a permissive agent
	if err != nil {
		robotAgent = c.permissiveAgent
	}

	c.robotAgents[u.Host] = robotAgent
	return robotAgent
}

func (c *client) createRobotAgent(u *url.URL) (robotAgent *robotstxt.Group, err error) {
	robotsURL := u.ResolveReference(robotsTxtPath)

	req, err := http.NewRequest("GET", robotsURL.String(), nil)
	if err != nil {
		return nil, err
	}

	// make sure request has our user agent
	req.Header.Add("User-Agent", c.userAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.Body != nil {
		defer func() {
			if cerr := resp.Body.Close(); err == nil && cerr != nil {
				err = cerr
			}
		}()
	}

	robotsData, err := robotstxt.FromResponse(resp)
	if err != nil {
		return nil, err
	}

	return robotsData.FindGroup(c.userAgent), nil
}
