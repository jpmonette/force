// Package force provides access to Salesforce various APIs
package force

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const (
	apiVersion = "v34.0"
	userAgent  = "github.com/jpmonette/force"
)

// Client is an HTTP client used to interact with the Salesforce API
type Client struct {
	client *http.Client

	// Base URL for API requests. Defaults to the production Salesforce API,
	// but can be set to a domain endpoint to use with Salesforce sandboxes.
	// BaseURL should always be specified with a trailing slash.
	BaseURL *url.URL

	// User agent used when communicating.
	UserAgent string
}

// NewClient returns a new Salesforce API client. If a nil httpClient is
// provided, http.DefaultClient will be used. To use API methods which require
// authentication, provide an http.Client that will perform the authentication
// for you (such as that provided by the golang.org/x/oauth2 library).
func NewClient(httpClient *http.Client, instanceUrl string) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, err := url.Parse(instanceUrl)

	if err != nil {
		return nil, err
	}

	c := &Client{client: httpClient, BaseURL: baseURL}

	return c, nil
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse("/services/data/" + apiVersion + urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}
	return req, nil
}
