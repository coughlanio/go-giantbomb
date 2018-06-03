# [WIP] go-giantbomb
[![Go Report Card](https://goreportcard.com/badge/github.com/coughlanio/go-giantbomb)](https://goreportcard.com/report/github.com/coughlanio/go-giantbomb)

<div style="text-align:center;width:120px"><img src ="https://static.giantbomb.com/uploads/original/14/148012/2649339-4392070099-22172.png" /></div>

This project was started as my own personal foray into Go development. Functionality is largely incomplete right now, with only a few endpoints implemented. Much of the groundwork is in place, so adding the remaining endpoints should be fairly trivial.
# Requirements

* Go > 1.10

# Installation

``` sh
go get -u github.com/coughlanio/go-giantbomb/giantbomb
```

# Examples

Several examples have been included to show off some basic functionality. You can run them using the following:

``` sh
GB_API_KEY=<Giant Bomb API Token> go run examples/game.go
```

# Usage

``` Golang
package main

import (
	"fmt"
	"go-giantbomb/giantbomb"
	"os"
)

func main() {

	client := giantbomb.NewClient(os.Getenv("GB_API_KEY"))

	games, err := client.Game.List(nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	guid := games.Results[0].GUID

	fmt.Printf("Fetching Game Info: %s\n", guid)

	game, err := client.Game.Get(guid, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Result: %#v\n", game)

}
```
