package force

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDescribeGlobal(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/services/data/v34.0/tooling/sobjects/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"encoding":"UTF-8","maxBatchSize":1,"sobjects":[{"activateable":false,"createable":true,"custom":false,"customSetting":false,"deletable":true,"deprecatedAndHidden":false,"feedEnabled":false,"keyPrefix":"01p","label":"Apex Class","labelPlural":"Apex Classes","layoutable":false,"mergeable":false,"name":"ApexClass","queryable":true,"replicateable":true,"retrieveable":true,"searchable":true,"triggerable":false,"undeletable":false,"updateable":true,"urls":{"rowTemplate":"/services/data/v34.0/tooling/sobjects/ApexClass/{ID}","describe":"/services/data/v34.0/tooling/sobjects/ApexClass/describe","sobject":"/services/data/v34.0/tooling/sobjects/ApexClass"}}]}`)
	})

	result, err := client.Tooling.DescribeGlobal()

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(result.SObjects))
	assert.Equal(t, "ApexClass", result.SObjects[0].Name)
}

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

func TestRunTests(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/services/data/v34.0/tooling/runTestsSynchronous/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"apexLogId":"","codeCoverage":[{"dmlInfo":[],"id":"01q20000000AmAWAA0","locationsNotCovered":[{"column":0,"line":2,"numExecutions":0,"time":-1},{"column":0,"line":5,"numExecutions":0,"time":-1},{"column":0,"line":6,"numExecutions":0,"time":-1}],"methodInfo":[],"name":"Events","namespace":"","numLocations":3,"numLocationsNotCovered":3,"soqlInfo":[],"soslInfo":[],"type":"Trigger"}],"codeCoverageWarnings":[{"id":"01q20000000AmAWAA0","message":"Test coverage of selected Apex Trigger is 0%, at least 1% test coverage is required","name":"Events","namespace":""},{"id":"01q20000000AmAWAA0","message":"Average test coverage across all Apex Classes and Triggers is 0%, at least 75% test coverage is required.","name":"","namespace":""}],"failures":[],"numFailures":0,"numTestsRun":1,"successes":[{"id":"01p20000004dwL0AAI","methodName":"test","name":"Tests_T","namespace":"","seeAllData":"","time":138}],"totalTime":140}`)
	})

	result, err := client.Tooling.RunTests([]string{"Tests_T"})

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(result.Successes))
	assert.Equal(t, "Tests_T", result.Successes[0].Name)
}

func TestRunTestsAsynchronous(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/services/data/v37.0/tooling/runTestsAsynchronous/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `"707E000004Btvxf"`)
	})

	result, err := client.Tooling.RunTestsAsynchronous(nil, nil, "", "RunLocalTests")

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, "707E000004Btvxf", result)
}

type SearchResponse struct {
	DurableID  string `json:"DurableId"`
	ID         string `json:"Id"`
	Attributes struct {
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"attributes"`
}

func TestSearch(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/services/data/v34.0/tooling/search/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[{"attributes":{"type":"EntityDefinition","url":"/services/data/v34.0/tooling/sobjects/EntityDefinition/4ie000000000005AAA"},"Id":"4ie000000000005AAA","DurableId":"AccountContactRole"},{"attributes":{"type":"EntityDefinition","url":"/services/data/v34.0/tooling/sobjects/EntityDefinition/4ie000000000006AAA"},"Id":"4ie000000000006AAA","DurableId":"AccountCleanInfo"}]`)
	})

	var result []SearchResponse

	err := client.Tooling.Search("FIND {account*} RETURNING EntityDefinition", &result)

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "AccountContactRole", result[0].DurableID)
}
