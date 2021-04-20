[![GoDoc](https://godoc.org/github.com/galaco/stringtable?status.svg)](https://godoc.org/github.com/galaco/stringtable)
[![Go report card](https://goreportcard.com/badge/github.com/galaco/stringtable)](hhttps://goreportcard.com/report/github.com/galaco/stringtable)
[![GolangCI](https://golangci.com/badges/github.com/galaco/stringtable.svg)](https://golangci.com/r/github.com/galaco/stringtable)
[![codecov](https://codecov.io/gh/galaco/stringtable/branch/master/graph/badge.svg)](https://codecov.io/gh/galaco/stringtable)
[![CircleCI](https://circleci.com/gh/galaco/stringtable.svg?style=svg)](https://circleci.com/gh/galaco/stringtable)

# Stringtable

> Stringtable is an indexed lookuptable containing 0 or more strings.


### Usage

Stringtable is a simple package. You can either:

* Create a new table
* Create a table from existing stringtable data (e.g. TexDataString* BSP lumps)

Here is a simple example:

```go

package main

import "github.com/galaco/stringtable"

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
