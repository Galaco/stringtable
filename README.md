[![GoDoc](https://godoc.org/github.com/golang-source-engine/stringtable?status.svg)](https://godoc.org/github.com/golang-source-engine/stringtable)
[![Go report card](https://goreportcard.com/badge/github.com/golang-source-engine/stringtable)](hhttps://goreportcard.com/report/github.com/golang-source-engine/stringtable)
[![GolangCI](https://golangci.com/badges/github.com/golang-source-engine/stringtable.svg)](https://golangci.com/r/github.com/golang-source-engine/stringtable)
[![codecov](https://codecov.io/gh/golang-source-engine/stringtable/branch/master/graph/badge.svg)](https://codecov.io/gh/golang-source-engine/stringtable)
[![CircleCI](https://circleci.com/gh/golang-source-engine/stringtable.svg?style=svg)](https://circleci.com/gh/golang-source-engine/stringtable)

# Stringtable

> Stringtable is an indexed lookuptable containing 0 or more strings.


### Usage

Stringtable is a simple package. You can either:

* Create a new table
* Create a table from existing stringtable data (e.g. TexDataString* BSP lumps)

Here is a simple example:

```go

package main

import "github.com/golang-source-engine/stringtable"

func main() {
    table := stringtable.New()
    val := "foo"
    index := table.AddString(val)

    s,err := table.FindString(index)
    if err != nil {
        panic(err)
    }
    if s != val {
        panic("returned string doesnt match stored")
    }
}
