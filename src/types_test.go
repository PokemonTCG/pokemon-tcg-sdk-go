package pokemontcgsdk

import (
	"sort"
	"testing"
)

func TestGetTypes(t *testing.T) {
	tests := []struct {
		types []string
	}{
		{
			[]string{"Colorless", "Darkness", "Dragon", "Fairy", "Fighting", "Fire", "Grass", "Lightning", "Metal", "Psychic", "Water"},
		},
	}

	for index, test := range tests {
		types, err := GetTypes()
		if err != nil {
			t.Fatalf("GetTypes failed: %s", err.Error())
		}

		if len(types) != len(test.types) {
			t.Errorf("Case %d: expected there to be %d items, but got %d", index, len(test.types), len(types))
		}

		//To ensure that it is sorted
		sort.Strings(types)

		for i, item := range test.types {
			if item != types[i] {
				t.Errorf("Case %d: expected %s, but got %s", index, item, types[i])
			}
		}
	}
}
