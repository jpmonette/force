# force

  force is a Go client library for accessing the [Salesforce API](https://developer.salesforce.com/docs/atlas.en-us.api_rest.meta/).

  [![GoDoc](https://godoc.org/github.com/jpmonette/force?status.svg)](https://godoc.org/github.com/jpmonette/force)
  [![TravisCI Build Status](https://travis-ci.org/jpmonette/force.svg)](https://travis-ci.org/jpmonette/force)
  [![CircleCI Build Status](https://circleci.com/gh/jpmonette/force.png?style=shield&circle-token=:circle-token)](https://circleci.com/gh/jpmonette/force)
  [![Coverage Status](https://coveralls.io/repos/jpmonette/force/badge.svg?branch=master&service=github)](https://coveralls.io/github/jpmonette/force?branch=master)

> ***WARNING:*** Both the documentation and the package itself is under heavy
> development and in a very early stage. That means, this repo is full of
> untested code and the API can break without any further notice. Therefore,
> it comes with absolutely no warranty at all. Feel free to browse or even
> contribute to it :)

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
