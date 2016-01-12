package force

import (
	"net/http"
	"net/http/httptest"
	"net/url"
)

var (
	mux    *http.ServeMux
	client *Client
	server *httptest.Server
)

func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// github client configured to use test server
	url, _ := url.Parse(server.URL)
	client, _ = NewClient(nil, url.String())
	client.BaseURL = url
}

// teardown closes the test HTTP server.
func teardown() {
	server.Close()
}
