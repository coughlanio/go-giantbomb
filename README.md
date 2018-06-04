# [WIP] go-giantbomb
[![Go Report Card](https://goreportcard.com/badge/github.com/coughlanio/go-giantbomb)](https://goreportcard.com/report/github.com/coughlanio/go-giantbomb)

<div style="text-align:center;width:120px"><img src ="https://static.giantbomb.com/uploads/original/14/148012/2649339-4392070099-22172.png" /></div>

This project was started as my own personal foray into Go development. Functionality is largely incomplete right now, with only a few endpoints implemented. Much of the groundwork is in place, so adding the remaining endpoints should be fairly trivial.
# Requirements

* Go > 1.10

# Authentication

You will need a Giant Bomb API token to use this library. You can get your token from [here.](https://www.giantbomb.com/api/)

# Installation

``` sh
go get -u github.com/coughlanio/go-giantbomb/giantbomb
```

# Tests

Tests have been included for the majority of the endpoints already implemented. You can run the test suite using the following:

``` sh
GB_API_KEY=<Giant Bomb API Token> go test ./...
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
