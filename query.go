// Package force provides access to Salesforce various APIs
package force

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
)

// Query is used for retrieving query performance feedback without executing
// the query
func (c *Client) Query(query string, v interface{}) (err error) {

	endpoint := fmt.Sprintf("/query/?q=%v", url.QueryEscape(query))
	req, err := c.NewRequest("GET", endpoint, nil)

	if err != nil {
		return
	}

	resp, err := c.client.Do(req)

	if err != nil {
		return
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}

	return
}
