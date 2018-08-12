package pokemontcgsdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	// EndPoint is the URL used to connect to the pokemontcg.io server. To view their docs
	// see https://docs.pokemontcg.io/
	EndPoint = "https://api.pokemontcg.io/v1/"
)

var (
	//EndPointCards used to obtain cards from the pokemontcg api
	EndPointCards = EndPoint + "/cards/"
)

//PokemonCard -- contains all the needed information about a pokemon card
type PokemonCard struct {
	ID                    string       `json:"id"`
	Name                  string       `json:name`
	NationalPokedexNumber int          `json:nationalPokedexNumber`
	ImageURL              string       `json:imageUrl`
	ImageURLHiRes         string       `json:imageUrlHiRes`
	Types                 []string     `json:types`
	SuperType             string       `json:supertype`
	SubType               string       `json:subtype`
	EvolvesFrom           string       `json:evolvesFrom`
	HP                    string       `json:hp`
	RetreatCost           []string     `json:retreatCost`
	Number                string       `json:number`
	Artist                string       `json:artist`
	Rarity                string       `json:rarity`
	Series                string       `json:series`
	Set                   string       `json:set`
	SetCode               string       `json:setCode`
	Attacks               []Attack     `json:attacks`
	Weaknesses            []Weakness   `json:weaknesses`
	Resistances           []Resistance `json:resistances`
	Ability               Ability      `json:ability`
	Text                  []string     `json:text`
}

// Attack holds information for pokemon attacks
type Attack struct {
	Cost                []string `json:cost`
	Name                string   `json:name`
	Text                string   `json:text`
	Damage              string   `json:damage`
	ConvertedEnergyCost int      `json:convertedEnergyCost`
}

// Weakness holds information for pokemon weaknesses
type Weakness struct {
	Type  string `json:type`
	Value string `json:type`
}

// Resistance holds information for pokemon weaknesses
type Resistance struct {
	Type  string `json:type`
	Value string `json:type`
}

// Ability holds pokemon ability information
type Ability struct {
	Name string `json:name`
	Text string `json:text`
}

// GetCardByID returns as single pokemon card.
func GetCardByID(ID string) (card PokemonCard, err error) {
	cards := make(map[string]PokemonCard)
	resp, err := http.Get(EndPointCards + ID)
	if err != nil {
		return card, fmt.Errorf("unable to http.Get card: %s", err.Error())
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&cards)
	if err != nil {
		return card, fmt.Errorf("Unable to decode json response: %s", err.Error())
	}

	var ok bool
	card, ok = cards["card"]
	if !ok {
		return card, fmt.Errorf("unable to get card from cards map: %s", err.Error())
	}
	return card, err

}
