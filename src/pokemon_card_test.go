package pokemontcgsdk

import (
	"strings"
	"testing"
)

/*
Test cards:
- https://api.pokemontcg.io/v1/cards/xy7-54
- https://api.pokemontcg.io/v1/cards/ex14-85
- https://api.pokemontcg.io/v1/cards/ecard3-132
- https://api.pokemontcg.io/v1/cards/xy7-74
- https://api.pokemontcg.io/v1/cards/xy7-53
- https://api.pokemontcg.io/v1/cards/xy7-52
*/
func TestGetCardByID(t *testing.T) {
	tests := []struct {
		cardID        string
		name          string
		imageURL      string
		imageURLHiRes string
		superType     string
		subType       string
		evolvesFrom   string
		hp            string
		rarity        string
		series        string
		set           string
		setCode       string
	}{
		{
			"xy7-54",
			"Gardevoir",
			"https://images.pokemontcg.io/xy7/54.png",
			"https://images.pokemontcg.io/xy7/54_hires.png",
			"Pokémon",
			"Stage 2",
			"Kirlia",
			"130",
			"Rare Holo",
			"XY",
			"Ancient Origins",
			"xy7",
		},
		{
			"ex14-85",
			"Windstorm",
			"https://images.pokemontcg.io/ex14/85.png",
			"https://images.pokemontcg.io/ex14/85_hires.png",
			"Trainer",
			"Item",
			"",     // evolvesFrom
			"None", // hp
			"Uncommon",
			"EX",
			"Crystal Guardians",
			"ex14",
		},
		{
			"ecard3-132",
			"Mirage Stadium",
			"https://images.pokemontcg.io/ecard3/132.png",
			"https://images.pokemontcg.io/ecard3/132_hires.png",
			"Trainer",
			"", // subtype
			"", // evlovesFrom
			"", // hp
			"Uncommon",
			"E-Card",
			"Skyridge",
			"ecard3",
		},
		{
			"xy7-74",
			"Forest of Giant Plants",
			"https://images.pokemontcg.io/xy7/74.png",
			"https://images.pokemontcg.io/xy7/74_hires.png",
			"Trainer",
			"Stadium",
			"", // evolvesFrom
			"", // hp
			"Uncommon",
			"XY",
			"Ancient Origins",
			"xy7",
		},
		{
			"xy7-53",
			"Kirlia",
			"https://images.pokemontcg.io/xy7/53.png",
			"https://images.pokemontcg.io/xy7/53_hires.png",
			"Pokémon",
			"Stage 1",
			"Ralts",
			"80",
			"Uncommon",
			"XY",
			"Ancient Origins",
			"xy7",
		},
		{
			"xy7-52",
			"Ralts",
			"https://images.pokemontcg.io/xy7/52.png",
			"https://images.pokemontcg.io/xy7/52_hires.png",
			"Pokémon",
			"Basic",
			"", // evolvesFrom
			"60",
			"Common",
			"XY",
			"Ancient Origins",
			"xy7",
		},
	}

	for index, test := range tests {
		card, err := GetCardByID(test.cardID)
		if err != nil {
			t.Fatalf("GetCardByID failed: %s", err.Error())
		}

		if !strings.EqualFold(card.Name, test.name) {
			t.Errorf("Case %d: Expected name to be %s, but got %s", index, test.name, card.Name)
		}

		if !strings.EqualFold(card.ImageURL, test.imageURL) {
			t.Errorf("Case %d: Expected %s, but got %s", index, test.imageURL, card.ImageURL)
		}

		if !strings.EqualFold(card.ImageURLHiRes, test.imageURLHiRes) {
			t.Errorf("Case %d: Expected %s, but got %s", index, test.imageURLHiRes, card.ImageURLHiRes)
		}

		if !strings.EqualFold(card.SuperType, test.superType) {
			t.Errorf("Case %d: Expected %s, but got %s", index, test.superType, card.SuperType)
		}

		if !strings.EqualFold(card.SubType, test.subType) {
			t.Errorf("Case %d: Expected %s, but got %s", index, test.subType, card.SubType)
		}

		if !strings.EqualFold(card.EvolvesFrom, test.evolvesFrom) {
			t.Errorf("Case %d: Expected %s, but got %s", index, test.evolvesFrom, card.EvolvesFrom)
		}

		if !strings.EqualFold(card.HP, test.hp) {
			t.Errorf("Case %d: Expected %s, but got %s", index, test.hp, card.HP)
		}
		if !strings.EqualFold(card.Rarity, test.rarity) {
			t.Errorf("Case %d: Expected %s, but got %s", index, test.rarity, card.Rarity)
		}
		if !strings.EqualFold(card.Series, test.series) {
			t.Errorf("Case %d: Expected %s, but got %s", index, test.series, card.Series)
		}
		if !strings.EqualFold(card.Set, test.set) {
			t.Errorf("Case %d: Expected %s, but got %s", index, test.set, card.Set)
		}
		if !strings.EqualFold(card.SetCode, test.setCode) {
			t.Errorf("Case %d: Expected %s, but got %s", index, test.setCode, card.SetCode)
		}

	}
}
