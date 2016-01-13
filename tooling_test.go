package force

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecuteAnonymous(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/services/data/v34.0/tooling/executeAnonymous/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"column":-1,"compileProblem":null,"compiled":true,"exceptionMessage":null,"exceptionStackTrace":null,"line":-1,"success":true}`)
	})

	result, err := client.Tooling.ExecuteAnonymous("System.debug('test');")

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, true, result.Success)
}

// QueryResponse is a Tooling API Query response object.
type QueryResponse struct {
	Records []struct {
		FullName string `json:"FullName"`
	} `json:"records"`
}

func TestQuery(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/services/data/v34.0/tooling/query/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"done":true,"entityTypeName":"Profile","queryLocator":"","records":[{"FullName":"Admin","attributes":{"type":"Profile","url":"/services/data/v34.0/tooling/sobjects/Profile/00N20000009gIZmEAM"}}],"size":1,"totalSize":1}`)
	})

	var result QueryResponse

	err := client.Tooling.Query("SELECT FullName FROM Profile", &result)

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, "Admin", result.Records[0].FullName)
}
