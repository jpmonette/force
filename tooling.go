package force

import "net/url"

// ToolingService handles communication with the Tooling API methods
// of the Salesforce Tooling API.
type ToolingService struct {
	client *Client
}

// Query executes a query against a Tooling API object and returns data that
// matches the specified criteria.
func (c *ToolingService) Query(soql string, v interface{}) (err error) {
	req, err := c.client.NewRequest("GET", "/tooling/query/?q="+url.QueryEscape(soql), nil)
	err = c.client.Do(req, &v)
	return
}
