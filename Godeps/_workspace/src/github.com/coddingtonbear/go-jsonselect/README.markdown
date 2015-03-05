
[![Build Status](https://travis-ci.org/coddingtonbear/go-jsonselect.png?branch=master)](https://travis-ci.org/coddingtonbear/go-jsonselect)
[![GoDoc](https://godoc.org/github.com/coddingtonbear/go-jsonselect?status.png)](http://godoc.org/github.com/coddingtonbear/go-jsonselect)

go-jsonselect
=============

A golang implementation of [JSONSelect](http://jsonselect.org/) modeled off of [@mwhooker's Python implementation](https://github.com/mwhooker/jsonselect)

Supports all levels of the JSONSelect specification.

Installation
------------

```
go get github.com/coddingtonbear/go-jsonselect
```

Usage
-----

[API documentation is available on Godoc](http://godoc.org/github.com/coddingtonbear/go-jsonselect).

JSONSelect offers many different selectors, and describing the full usage
is out of the scope of this document, but should you need an overview, 
consult the [JSONSelect documentation](http://jsonselect.org/#docs).

Example
-------

Selecting all objects having a rating greater than 70 that are subordinate to a key "beers":

```golang
package main

import (
    "fmt"
    "github.com/coddingtonbear/go-jsonselect"
)

var json string = `
    {
        "beers": [
            {
                "title": "alpha",
                "rating": 50
            },
            {
                "title": "beta",
                "rating": 90
            }
        ]
    }
`

func main() {
    parser, _ := jsonselect.CreateParserFromString(json)
    results, _ := parser.GetValues(".beers object:has(.rating:expr(x>70))")
    fmt.Print(results)
    // Results [map[title: beta rating: 90]]
}
```

