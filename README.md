[![OpenSSF Scorecard](https://api.securityscorecards.dev/projects/github.com/bodgit/xraylogr/badge)](https://securityscorecards.dev/viewer/?uri=github.com/bodgit/xraylogr)
[![OpenSSF Best Practices](https://www.bestpractices.dev/projects/10667/badge)](https://www.bestpractices.dev/projects/10667)
[![GitHub release](https://img.shields.io/github/v/release/bodgit/xraylogr)](https://github.com/bodgit/xraylogr/releases)
[![Build Status](https://img.shields.io/github/actions/workflow/status/bodgit/xraylogr/build.yml?branch=main)](https://github.com/bodgit/xraylogr/actions?query=workflow%3ABuild)
[![Coverage Status](https://coveralls.io/repos/github/bodgit/xraylogr/badge.svg?branch=main)](https://coveralls.io/github/bodgit/xraylogr?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/bodgit/xraylogr)](https://goreportcard.com/report/github.com/bodgit/xraylogr)
[![GoDoc](https://godoc.org/github.com/bodgit/xraylogr?status.svg)](https://godoc.org/github.com/bodgit/xraylogr)
![Go version](https://img.shields.io/badge/Go-1.24-brightgreen.svg)
![Go version](https://img.shields.io/badge/Go-1.23-brightgreen.svg)

# xraylogr

This simple package provides a means to use [github.com/go-logr/logr](https://pkg.go.dev/github.com/go-logr/logr) with the [github.com/aws/aws-xray-sdk-go](https://pkg.go.dev/github.com/aws/aws-xray-sdk-go) AWS X-Ray tracing SDK.

An example:

```golang
l := stdr.New(nil)

nl, err := xraylogr.New(l)
if err != nil {
	log.Fatal(err)
}

xray.SetLogger(nl)
```
