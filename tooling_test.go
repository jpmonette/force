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
