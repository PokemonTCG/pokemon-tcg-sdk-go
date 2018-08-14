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
