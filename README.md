Intro
=====

An implementation of Connect Four in Go. I typically use this
implementation to get used to testing and simple language semantics when
learning a new language.

[![Build Status](https://travis-ci.org/jamiely/connect-four-go.png?branch=master)](https://travis-ci.org/jamiely/connect-four-go)

Usage
=====

If your `GOPATH` is not set, then make sure to set it to something like
`$HOME/gopath`.

To run use the package, create a go program `run.go` like:

```go
package main

import (
  . "github.com/jamiely/connect-four-go"
)

func main() {
  ConsoleUI()
}
```

Install the package using `go get github.com/jamiely/connect-four-go`.

Then, run the go application via: `go run run.go`

To test the application, execute:

```bash
cd src/connect_four
GOPATH=`pwd` go get github.com/orfjackal/gospec/src/gospec
GOPATH=`pwd` go test
```

