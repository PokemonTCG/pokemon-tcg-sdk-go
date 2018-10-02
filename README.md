# pokemon-tcg-sdk-go
## PokemonTCG Golang SDK

[![Build Status](https://api.travis-ci.org/PokemonTCG/pokemon-tcg-sdk-go.svg?branch=master)](https://travis-ci.org/PokemonTCG/pokemon-tcg-sdk-go)    [![Go Report Card](https://goreportcard.com/badge/github.com/PokemonTCG/pokemon-tcg-sdk-go)](https://goreportcard.com/report/github.com/PokemonTCG/pokemon-tcg-sdk-go)    [![GoDoc](https://godoc.org/github.com/PokemonTCG/pokemon-tcg-sdk-go/src?status.svg)](https://godoc.org/github.com/PokemonTCG/pokemon-tcg-sdk-go/src)    [![Coverage Status](https://coveralls.io/repos/github/PokemonTCG/pokemon-tcg-sdk-go/badge.svg?branch=master)](https://coveralls.io/github/PokemonTCG/pokemon-tcg-sdk-go?branch=master)

### Install
`go get github.com/PokemonTCG/pokemon-tcg-sdk-go/src`

### Basic Usage

	package main

	import (
		"fmt"
		"log"

		"github.com/PokemonTCG/pokemon-tcg-sdk-go/src"
	)

	func main() {
		noParams := make(map[string]string)
		cards, err := pokemontcgsdk.GetCards(noParams)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(cards)
	}

See [examples](https://github.com/PokemonTCG/pokemon-tcg-sdk-go/tree/master/_examples) section for more.

### API Documentation

See the [pokemontcg website](https://pokemontcg.io/) ([The docs](https://docs.pokemontcg.io/)).


