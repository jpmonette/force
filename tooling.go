package force

import "net/url"

// ToolingService handles communication with the Tooling API methods
// of the Salesforce Tooling API.
type ToolingService struct {
	client *Client
}

// ExecuteAnonymousResult specifies information about whether or not the
// compile and run of the code was successful
type ExecuteAnonymousResult struct {
	Column              int    `json:"column"`
	CompileProblem      string `json:"compileProblem"`
	Compiled            bool   `json:"compiled"`
	ExceptionMessage    string `json:"exceptionMessage"`
	ExceptionStackTrace string `json:"exceptionStackTrace"`
	Line                int    `json:"line"`
	Success             bool   `json:"success"`
}

// ExecuteAnonymous executes the specified block of Apex anonymously and
// returns the result.
func (c *ToolingService) ExecuteAnonymous(apex string) (result ExecuteAnonymousResult, err error) {
	req, err := c.client.NewRequest("GET", "/tooling/executeAnonymous/?anonymousBody="+url.QueryEscape(apex), nil)
	err = c.client.Do(req, &result)
	return
}

// Query executes a query against a Tooling API object and returns data that
// matches the specified criteria.
func (c *ToolingService) Query(soql string, v interface{}) (err error) {
	req, err := c.client.NewRequest("GET", "/tooling/query/?q="+url.QueryEscape(soql), nil)
	err = c.client.Do(req, &v)
	return
}
