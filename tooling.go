package force

import (
	"net/url"
	"strings"
)

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

// RunTestsResult contains information about the execution of unit tests,
// including whether unit tests were completed successfully, code coverage
// results, and failures.
type RunTestsResult struct {
	ApexLogID            string                `json:"apexLogId"`
	CodeCoverage         []CodeCoverageResult  `json:"codeCoverage"`
	CodeCoverageWarnings []CodeCoverageWarning `json:"codeCoverageWarnings"`
	Failures             []RunTestFailure      `json:"failures"`
	NumFailures          int                   `json:"numFailures"`
	NumTestsRun          int                   `json:"numTestsRun"`
	Successes            []RunTestSuccess      `json:"successes"`
	TotalTime            float64               `json:"totalTime"`
}

// CodeCoverageResult contains the details of the code coverage for the
// specified unit tests.
type CodeCoverageResult struct {
	DmlInfo                []CodeLocation `json:"dmlInfo"`
	ID                     string         `json:"id"`
	LocationsNotCovered    []CodeLocation `json:"locationsNotCovered"`
	MethodInfo             []CodeLocation `json:"methodInfo"`
	Name                   string         `json:"name"`
	Namespace              string         `json:"namespace"`
	NumLocations           int            `json:"numLocations"`
	NumLocationsNotCovered int            `json:"numLocationsNotCovered"`
	SoqlInfo               []CodeLocation `json:"soqlInfo"`
	SoslInfo               []CodeLocation `json:"soslInfo"`
	Type                   string         `json:"type"`
}

// CodeCoverageWarning contains results include both the total number of lines
// that could have been executed, as well as the number, line, and column
// positions of code that was not executed.
type CodeCoverageWarning struct {
	Id        string `json:"id"`
	Message   string `json:"message"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

// RunTestFailure contains information about failures during the unit test run.
type RunTestFailure struct {
	Id         string  `json:"id"`
	Message    string  `json:"message"`
	MethodName string  `json:"methodName"`
	Name       string  `json:"name"`
	Namespace  string  `json:"namespace"`
	SeeAllData string  `json:"seeAllData"`
	StackTrace string  `json:"stackTrace"`
	Time       float64 `json:"time"`
	Type       string  `json:"string"`
}

// RunTestSuccess contains information about successes during the unit test run.
type RunTestSuccess struct {
	Id         string  `json:"id"`
	MethodName string  `json:"methodName"`
	Name       string  `json:"name"`
	Namespace  string  `json:"namespace"`
	SeeAllData string  `json:"seeAllData"`
	Time       float64 `json:"time"`
}

// CodeLocation contains if any code is not covered, the line and column of the
// code not tested, and the number of times the code was executed
type CodeLocation struct {
	Column        int     `json:"column"`
	Line          int     `json:"line"`
	NumExecutions int     `json:"numExecutions"`
	Time          float64 `json:"time"`
}

// RunTests executes the tests in the specified classes using the synchronous
// test execution mechanism.
func (c *ToolingService) RunTests(classnames []string) (result RunTestsResult, err error) {
	req, err := c.client.NewRequest("GET", "/tooling/runTestsSynchronous/?classnames="+url.QueryEscape(strings.Join(classnames, ",")), nil)
	err = c.client.Do(req, &result)
	return
}
