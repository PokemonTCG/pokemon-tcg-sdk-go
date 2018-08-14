# go-pokemontcgsdk
## Unofficial PokemonTCG Golang SDK

[![Build Status](https://api.travis-ci.org/crazcalm/go-pokemontcgsdk.svg?branch=master)](https://travis-ci.org/crazcalm/term-quiz)    [![Go Report Card](https://goreportcard.com/badge/github.com/crazcalm/go-pokemontcgsdk)](https://goreportcard.com/report/github.com/crazcalm/go-pokemontcgsdk)    [![GoDoc](https://godoc.org/github.com/crazcalm/go-pokemontcgsdk/src?status.svg)](https://godoc.org/github.com/crazcalm/go-pokemontcgsdk/src)    [![Coverage Status](https://coveralls.io/repos/github/crazcalm/go-pokemontcgsdk/badge.svg?branch=master)](https://coveralls.io/github/crazcalm/go-pokemontcgsdk?branch=master)

### Install
`go get github.com/crazcalm/go-pokemontcgsdk/src`

### Basic Usage

	package main

	import (
		"fmt"
		"log"

		"github.com/crazcalm/go-pokemontcgsdk/src"
	)

	func main() {
		noParams := make(map[string]string)
		cards, err := pokemontcgsdk.GetCards(noParams)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(cards)
	}

See [examples](https://github.com/crazcalm/go-pokemontcgsdk/tree/master/_examples) section for more.

### API Documentation

See the [pokemontcg website](https://pokemontcg.io/) ([The docs](https://docs.pokemontcg.io/)).


