# force

force is a Go client library for accessing the [Salesforce API](https://developer.salesforce.com/docs/atlas.en-us.api_rest.meta/).

[![Build Status](https://travis-ci.org/jpmonette/force.svg)](https://travis-ci.org/jpmonette/force)

**Documentation:** [![GoDoc](https://godoc.org/github.com/jpmonette/force?status.svg)](https://godoc.org/github.com/jpmonette/force)  

## Usage

```go
import "github.com/jpmonette/force"
```

Construct a new Force client, then use the various services on the client to access different parts of the Salesforce API. For example, to retrieve query performance feedback:

```go
  c, _ := force.NewClient(client, "http://emea.salesforce.com/")
  explain, err := c.QueryExplain("SELECT Id, Name, OwnerId FROM Account LIMIT 10")
```

### Authentication

The force library does not directly handle authentication.  Instead, when
creating a new client, pass an `http.Client` that can handle authentication for
you.  The easiest and recommended way to do this is using the
[oauth2](https://godoc.org/golang.org/x/oauth2) library, but you can always use
any other library that provides an `http.Client`.


## Roadmap

This library is being initially developed for one of my internal project,
so API methods will likely be implemented in the order that they are
needed by my project. Eventually, I would like to cover the entire
Salesforce API, so contributions are of course [always welcome][contributing].  The
calling pattern is pretty well established, so adding new methods is relatively
straightforward.

[contributing]: CONTRIBUTING.md


## License

This library is distributed under the MIT license found in the [LICENSE](./LICENSE)
file.
