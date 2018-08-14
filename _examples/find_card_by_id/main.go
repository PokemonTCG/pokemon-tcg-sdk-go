package main

import (
	"fmt"
	"log"

	"github.com/crazcalm/go-pokemontcgsdk/src"
)

func main() {
	card, err := pokemontcgsdk.GetCardByID("xy7-54")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(card)
}
