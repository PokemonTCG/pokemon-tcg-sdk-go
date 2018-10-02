package main

import (
	"fmt"
	"log"

	"github.com/PokemonTCG/pokemon-tcg-sdk-go/src"
)

func main() {
	card, err := pokemontcgsdk.GetCardByID("xy7-54")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(card)
}
