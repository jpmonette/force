// Package force provides access to Salesforce various APIs
package force

import (
	"fmt"
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

	err = c.Do(req, &v)

	if err != nil {
		return
	}

	return
}

// QueryExplain is used for retrieving query performance feedback without
// executing the query
func (c *Client) QueryExplain(query string) (explain QueryExplainResponse, err error) {

	endpoint := fmt.Sprintf("/query/?explain=%v", url.QueryEscape(query))
	req, err := c.NewRequest("GET", endpoint, nil)

	if err != nil {
		return
	}

	err = c.Do(req, &explain)

	if err != nil {
		return
	}

	return
}

// QueryExplainResponse is returned by QueryExplain
type QueryExplainResponse struct {
	Plans []struct {
		Cardinality          int      `json:"cardinality"`
		Fields               []string `json:"fields"`
		LeadingOperationType string   `json:"leadingOperationType"`
		RelativeCost         float64  `json:"relativeCost"`
		SobjectCardinality   int      `json:"sobjectCardinality"`
		SobjectType          string   `json:"sobjectType"`
		Notes                []struct {
			Description   string   `json:"description"`
			Fields        []string `json:"fields"`
			TableEnumOrID string   `json:"tableEnumOrId"`
		} `json:"notes"`
	} `json:"plans"`
}
