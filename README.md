Intro
=====

An implementation of Connect Four in Go. I typically use this
implementation to get used to testing and simple language semantics when
learning a new language.

[![Build Status](https://travis-ci.org/jamiely/connect-four-go.png?branch=master)](https://travis-ci.org/jamiely/connect-four-go)

Usage
=====

To run the application, execute:

```bash
GOPATH=`pwd` go run run.go
```

To test the application, execute:

```bash
cd src/connect_four
GOPATH=`pwd` go get github.com/orfjackal/gospec/src/gospec
GOPATH=`pwd` go test
```

