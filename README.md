Go einfo [![Last release](https://img.shields.io/github/release/euskadi31/go-einfo.svg)](https://github.com/euskadi31/go-einfo/releases/latest) [![Documentation](https://godoc.org/github.com/euskadi31/go-einfo?status.svg)](https://godoc.org/github.com/euskadi31/go-einfo)
=========

[![Go Report Card](https://goreportcard.com/badge/github.com/euskadi31/go-einfo)](https://goreportcard.com/report/github.com/euskadi31/go-einfo)

| Branch  | Status | Coverage |
|---------|--------|----------|
| master  | [![Build Status](https://img.shields.io/travis/euskadi31/go-einfo/master.svg)](https://travis-ci.org/euskadi31/go-einfo) | [![Coveralls](https://img.shields.io/coveralls/euskadi31/go-einfo/master.svg)](https://coveralls.io/github/euskadi31/go-einfo?branch=master) |


~~~go
package main

import einfo "github.com/euskadi31/go-einfo"

func main() {
    einfo.Info("test info")
    einfo.Warn("test warn")
    einfo.Error("test error")

    einfo.Begin("Fetch data")
    einfo.End(true, "Cannot fetch data")

    einfo.Begin("Fetch data")
    einfo.End(false, "Cannot fetch data")
}
~~~

![demo](files/demo.png?raw=true "Demo")
