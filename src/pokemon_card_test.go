package pokemontcgsdk

import (
	"sort"
	"strings"
	"testing"
)

func TestFormatQuery(t *testing.T) {
	tests := []struct {
		params map[string]string
		expect string
	}{
		{map[string]string{}, EndPointCards},
		{map[string]string{"supertype": "pokemon", "types": "dragon|fire|flying", "hp": "gt100"}, "https://api.pokemontcg.io/v1/cards?hp=gt100&supertype=pokemon&types=dragon%7Cfire%7Cflying"},
		{map[string]string{"page": "50", "pageSize": "500"}, "https://api.pokemontcg.io/v1/cards?page=50&pageSize=500"},
	}

	for index, test := range tests {
		result, err := formatQuery(EndPointCards, test.params)
		if err != nil {
			t.Fatalf("formatQuery failed: %s", err.Error())
		}

		if !strings.EqualFold(result, test.expect) {
			t.Errorf("Case %d: Expected %s, but got %s", index, test.expect, result)
		}
	}
}

func TestGetCardByIDBadID(t *testing.T) {
	expectedError := "Unable to decode json response: json: cannot unmarshal number into Go value of type pokemontcgsdk.PokemonCard"
	_, err := GetCardByID("BadInput")
	if !strings.EqualFold(expectedError, err.Error()) {
		t.Errorf("Expected error to say '%s',but got '%s'", expectedError, err.Error())
	}
}

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
		cardID                string
		name                  string
		nationalPokedexNumber int
		imageURL              string
		imageURLHiRes         string
		types                 []string
		superType             string
		subType               string
		evolvesFrom           string
		ability               Ability
		hp                    string
		retreatCost           []string
		convertedRetreatCost  int
		number                string
		artist                string
		rarity                string
		series                string
		set                   string
		setCode               string
		attacks               []Attack
		resistances           []Resistance
		weaknesses            []Weakness
	}{
		{
			"xy7-54",
			"Gardevoir",
			282,
			"https://images.pokemontcg.io/xy7/54.png",
			"https://images.pokemontcg.io/xy7/54_hires.png",
			[]string{"Fairy"},
			"Pokémon",
			"Stage 2",
			"Kirlia",
			Ability{
				Name: "Bright Heal",
				Text: "Once during your turn (before your attack), you may heal 20 damage from each of your Pokémon.",
				Type: "Ability"},
			"130",
			[]string{"Colorless", "Colorless"},
			2,
			"54",
			"TOKIYA",
			"Rare Holo",
			"XY",
			"Ancient Origins",
			"xy7",
			[]Attack{
				Attack{
					Cost:                []string{"Colorless", "Colorless", "Colorless"},
					Name:                "Telekinesis",
					Damage:              "",
					Text:                "This attack does 50 damage to 1 of your opponent's Pokémon. This attack's damage isn't affected by Weakness or Resistance.",
					ConvertedEnergyCost: 3}},
			[]Resistance{Resistance{Type: "Darkness", Value: "-20"}},
			[]Weakness{Weakness{Type: "Metal", Value: "×2"}},
		},
		{
			"ex14-85",
			"Windstorm",
			0, // national pokendex number
			"https://images.pokemontcg.io/ex14/85.png",
			"https://images.pokemontcg.io/ex14/85_hires.png",
			[]string{}, // types
			"Trainer",
			"Item",
			"", // evolvesFrom
			Ability{Name: "", Text: "", Type: ""},
			"None",     // hp
			[]string{}, // retreatCost
			0,          // convertedRetreatCost
			"85",
			"Ryo Ueda",
			"Uncommon",
			"EX",
			"Crystal Guardians",
			"ex14",
			[]Attack{},
			[]Resistance{},
			[]Weakness{},
		},
		{
			"ecard3-132",
			"Mirage Stadium",
			0,
			"https://images.pokemontcg.io/ecard3/132.png",
			"https://images.pokemontcg.io/ecard3/132_hires.png",
			[]string{}, // types
			"Trainer",
			"", // subtype
			"", // evlovesFrom
			Ability{Name: "", Text: "", Type: ""},
			"",         // hp
			[]string{}, // retreatCost
			0,          // covertedRetreatCost
			"132",
			"Midori Harada",
			"Uncommon",
			"E-Card",
			"Skyridge",
			"ecard3",
			[]Attack{},
			[]Resistance{},
			[]Weakness{},
		},
		{
			"xy7-74",
			"Forest of Giant Plants",
			0, // national pokedex number
			"https://images.pokemontcg.io/xy7/74.png",
			"https://images.pokemontcg.io/xy7/74_hires.png",
			[]string{}, // types
			"Trainer",
			"Stadium",
			"", // evolvesFrom
			Ability{Name: "", Text: "", Type: ""},
			"",         // hp
			[]string{}, // retreatCost
			0,          // covertedRetreatCost
			"74",
			"Ryo Ueda",
			"Uncommon",
			"XY",
			"Ancient Origins",
			"xy7",
			[]Attack{},
			[]Resistance{},
			[]Weakness{},
		},
		{
			"xy7-53",
			"Kirlia",
			281,
			"https://images.pokemontcg.io/xy7/53.png",
			"https://images.pokemontcg.io/xy7/53_hires.png",
			[]string{"Fairy"},
			"Pokémon",
			"Stage 1",
			"Ralts",
			Ability{Name: "", Text: "", Type: ""},
			"80",
			[]string{"Colorless"},
			1,
			"53",
			"match",
			"Uncommon",
			"XY",
			"Ancient Origins",
			"xy7",
			[]Attack{
				Attack{
					Cost:                []string{"Colorless"},
					Name:                "Calm Mind",
					Damage:              "",
					Text:                "Heal 30 damage from this Pokémon.",
					ConvertedEnergyCost: 1},
				Attack{
					Cost:                []string{"Fairy", "Colorless", "Colorless"},
					Name:                "Magical Shot",
					Damage:              "50",
					Text:                "",
					ConvertedEnergyCost: 3}},
			[]Resistance{Resistance{Type: "Darkness", Value: "-20"}},
			[]Weakness{Weakness{Type: "Metal", Value: "×2"}},
		},
		{
			"xy7-52",
			"Ralts",
			280,
			"https://images.pokemontcg.io/xy7/52.png",
			"https://images.pokemontcg.io/xy7/52_hires.png",
			[]string{"Fairy"},
			"Pokémon",
			"Basic",
			"", // evolvesFrom
			Ability{Name: "", Text: "", Type: ""},
			"60",
			[]string{"Colorless"},
			1,
			"52",
			"Aya Kusube",
			"Common",
			"XY",
			"Ancient Origins",
			"xy7",
			[]Attack{
				Attack{
					Cost:                []string{"Colorless"},
					Name:                "Mumble",
					Damage:              "10",
					Text:                "",
					ConvertedEnergyCost: 1},
				Attack{
					Cost:                []string{"Fairy", "Colorless"},
					Name:                "Magical Shot",
					Damage:              "20",
					Text:                "",
					ConvertedEnergyCost: 2}},
			[]Resistance{Resistance{Type: "Darkness", Value: "-20"}},
			[]Weakness{Weakness{Type: "Metal", Value: "×2"}},
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

		if card.NationalPokedexNumber != test.nationalPokedexNumber {
			t.Errorf("Case %d: Expected national pokedex number to be %d, but got %d", index, test.nationalPokedexNumber, card.NationalPokedexNumber)
		}

		if !strings.EqualFold(card.ImageURL, test.imageURL) {
			t.Errorf("Case %d: Expected %s, but got %s", index, test.imageURL, card.ImageURL)
		}

		if !strings.EqualFold(card.ImageURLHiRes, test.imageURLHiRes) {
			t.Errorf("Case %d: Expected %s, but got %s", index, test.imageURLHiRes, card.ImageURLHiRes)
		}

		if len(card.Types) != len(test.types) {
			t.Errorf("Case %d: Expected %v, but got %v", index, test.types, card.Types)
		}

		sort.Strings(card.Types)
		sort.Strings(test.types)
		for i, item := range test.types {
			if !strings.EqualFold(item, card.Types[i]) {
				t.Logf("Case %d: Expected %v, but got %v", index, test.types, card.Types)
				t.Errorf("--> Item %d: expected %s, but got %s", i, item, card.Types[i])
			}
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

		if !strings.EqualFold(card.Ability.Name, test.ability.Name) {
			t.Errorf("Case %d: Expected %s, but got %s", index, test.ability.Name, card.Ability.Name)
		}

		if !strings.EqualFold(card.Ability.Text, test.ability.Text) {
			t.Errorf("Case %d: Expected %s, but got %s", index, test.ability.Text, card.Ability.Text)
		}

		if !strings.EqualFold(card.Ability.Type, test.ability.Type) {
			t.Errorf("Case %d: Expected %s, but got %s", index, test.ability.Type, card.Ability.Type)
		}

		if !strings.EqualFold(card.HP, test.hp) {
			t.Errorf("Case %d: Expected %s, but got %s", index, test.hp, card.HP)
		}

		if len(card.RetreatCost) != len(test.retreatCost) {
			t.Errorf("Case %d: Expected %v, but got %v", index, test.retreatCost, card.RetreatCost)
		}

		sort.Strings(card.RetreatCost)
		sort.Strings(test.retreatCost)
		for i, item := range test.retreatCost {
			if !strings.EqualFold(item, card.RetreatCost[i]) {
				t.Logf("Case %d: Expected %v, but got %v", index, test.retreatCost, card.RetreatCost)
				t.Errorf("--> Item %d: expected %s, but got %s", i, item, card.RetreatCost[i])
			}
		}

		if card.ConvertedRetreatCost != test.convertedRetreatCost {
			t.Errorf("Case %d: Expected %d, but got %d", index, test.convertedRetreatCost, card.ConvertedRetreatCost)
		}

		if !strings.EqualFold(card.Name, test.name) {
			t.Errorf("Case %d: Expected %s, but got %s", index, test.rarity, card.Rarity)
		}

		if !strings.EqualFold(card.Artist, test.artist) {
			t.Errorf("Case %d: Expected %s, but got %s", index, test.rarity, card.Rarity)
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

		for i := 0; i < len(test.attacks); i++ {
			for count, energy := range test.attacks[i].Cost {
				if !strings.EqualFold(energy, card.Attacks[i].Cost[count]) {
					t.Errorf("Case %d: Expected %v, but got %v", index, test.attacks[i].Cost, card.Attacks[i].Cost)
				}
			}

			if card.Attacks[i].ConvertedEnergyCost != test.attacks[i].ConvertedEnergyCost {
				t.Errorf("Case %d: Expected %d, but got %d", index, test.attacks[i].ConvertedEnergyCost, card.Attacks[i].ConvertedEnergyCost)
			}

			if !strings.EqualFold(card.Attacks[i].Name, test.attacks[i].Name) {
				t.Errorf("Case %d: Expected %s, but got %s", index, test.attacks[i].Name, card.Attacks[i].Name)
			}

			if !strings.EqualFold(card.Attacks[i].Text, test.attacks[i].Text) {
				t.Errorf("Case %d: Expected %s, but got %s", index, test.attacks[i].Text, card.Attacks[i].Text)
			}

			if !strings.EqualFold(card.Attacks[i].Damage, test.attacks[i].Damage) {
				t.Errorf("Case %d: Expected %s, but got %s", index, test.attacks[i].Damage, card.Attacks[i].Damage)
			}
		}

		for i := 0; i < len(test.resistances); i++ {
			if !strings.EqualFold(card.Resistances[i].Type, test.resistances[i].Type) {
				t.Errorf("Case %d: Expected %s, but got %s", index, test.resistances[i].Type, card.Resistances[i].Type)
			}

			if !strings.EqualFold(card.Resistances[i].Value, test.resistances[i].Value) {
				t.Errorf("Case %d: Expected %s, but got %s", index, test.resistances[i].Value, card.Resistances[i].Value)
			}
		}

		for i := 0; i < len(test.weaknesses); i++ {
			if !strings.EqualFold(card.Weaknesses[i].Type, test.weaknesses[i].Type) {
				t.Errorf("Case %d: Expected %s, but got %s", index, test.weaknesses[i].Type, card.Weaknesses[i].Type)
			}

			if !strings.EqualFold(card.Weaknesses[i].Value, test.weaknesses[i].Value) {
				t.Errorf("Case %d: Expected %s, but got %s", index, test.weaknesses[i].Value, card.Weaknesses[i].Value)
			}
		}

	}
}

/*
Test cards:
- https://api.pokemontcg.io/v1/cards/xy7-54
- https://api.pokemontcg.io/v1/cards/ex14-85
- https://api.pokemontcg.io/v1/cards/ecard3-132
- https://api.pokemontcg.io/v1/cards/xy7-74
- https://api.pokemontcg.io/v1/cards/xy7-53
- https://api.pokemontcg.io/v1/cards/xy7-52
*/
func TestGetCard(t *testing.T) {
	tests := []struct {
		params        map[string]string
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
			map[string]string{"id": "xy7-54"},
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
			map[string]string{"id": "ex14-85"},
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
			map[string]string{"id": "ecard3-132"},
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
			map[string]string{"id": "xy7-74"},
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
			map[string]string{"id": "xy7-53"},
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
			map[string]string{"id": "xy7-52"},
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
		cards, err := GetCards(test.params)
		card := cards[0]
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
