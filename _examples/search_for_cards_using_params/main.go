package main

import (
	"fmt"
	"log"

	"github.com/PokemonTCG/pokemon-tcg-sdk-go/src"
)

func main() {
	params := map[string]string{"supertype": "pokemon", "types": "dragon|fire|flying", "hp": "gt100"}
	cards, err := pokemontcgsdk.GetCards(params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cards)
}
