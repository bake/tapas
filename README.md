# Tapas API client

[![GoDoc](https://godoc.org/github.com/bake/tapas?status.svg)](https://pkg.go.dev/github.com/bake/tapas)
[![Go Report Card](https://goreportcard.com/badge/github.com/bake/tapas)](https://goreportcard.com/report/github.com/bake/tapas)

A Go client for Tapas API at api.tapas.io.

```go
package main

import (
	"fmt"
	"log"

	"github.com/bake/tapas"
)

func main() {
	t := tapas.New()
	_, cs, err := t.Search("sarah andersen")
	if err != nil {
		log.Fatal(err)
	}
	s, err := t.Series(cs[0].ID)
	if err != nil {
		log.Fatal(err)
	}
	es, err := t.Episodes(cs[0].ID)
	if err != nil {
		log.Fatal(err)
	}
	e, err := t.Episode(cs[0].ID, es[0].ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s - %q by %s\n", s, e, s.Creators[0])
	for _, img := range e.Contents {
		fmt.Println(img)
	}
	// Sarah's Scribbles - "I Should Be vs I Am" by Sarah Andersen
	// https://....jpg
}
```
