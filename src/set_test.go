package pokemontcgsdk

import (
	"strings"
	"testing"
)

func TestGetSetByIDBadID(t *testing.T) {
	expectedError := "Unable to decode json response: json: cannot unmarshal number into Go value of type pokemontcgsdk.Set"
	_, err := GetSetByID("BadInput")
	if !strings.EqualFold(expectedError, err.Error()) {
		t.Errorf("Expected error to say '%s',but got '%s'", expectedError, err.Error())
	}
}

func TestGetSets(t *testing.T) {
	tests := []struct {
		params            map[string]string
		expectedNumOfSets int
	}{
		{
			map[string]string{"series": "base"},
			7,
		},
		{
			map[string]string{"series": "xy"},
			16,
		},
	}

	for index, test := range tests {
		sets, err := GetSets(test.params)
		if err != nil {
			t.Fatalf("GetSets failed: %s", err.Error())
		}

		if len(sets) != test.expectedNumOfSets {
			t.Errorf("Case %d: Expected %d sets, but got %d", index, len(sets), test.expectedNumOfSets)
		}

	}
}

/*
Test sets:
- https://api.pokemontcg.io/v1/sets/xy1
- https://api.pokemontcg.io/v1/sets/base1
*/
func TestGetSetByID(t *testing.T) {
	tests := []struct {
		code          string
		ptcgoCode     string
		name          string
		series        string
		totalCards    int
		standardLegal bool
		expandedLegal bool
		symbolURL     string
		logoURL       string
		releasedDate  string
		updatedAt     string
	}{
		{
			"base1",
			"BS",
			"Base",
			"Base",
			102,
			false,
			false,
			"https://images.pokemontcg.io/base1/symbol.png",
			"https://images.pokemontcg.io/base1/logo.png",
			"",
			"\"06/06/2018 20:35:00\"",
		},
		{
			"xy1",
			"XY",
			"XY",
			"XY",
			146,
			false,
			true,
			"https://images.pokemontcg.io/xy1/symbol.png",
			"https://images.pokemontcg.io/xy1/logo.png",
			"",
			"\"03/04/2018 10:35:00\"",
		},
	}

	for index, test := range tests {
		set, err := GetSetByID(test.code)
		if err != nil {
			t.Fatalf("GetSetBtID failed: %s", err.Error())
		}

		if !strings.EqualFold(test.code, set.Code) {
			t.Errorf("Case %d: Expected %v, but got %v", index, test.code, set.Code)
		}
		if !strings.EqualFold(test.ptcgoCode, set.PtcgoCode) {
			t.Errorf("Case %d: Expected %v, but got %v", index, test.ptcgoCode, set.PtcgoCode)
		}
		if !strings.EqualFold(test.name, set.Name) {
			t.Errorf("Case %d: Expected %v, but got %v", index, test.name, set.Name)
		}
		if !strings.EqualFold(test.series, set.Series) {
			t.Errorf("Case %d: Expected %v, but got %v", index, test.series, set.Series)
		}
		if test.totalCards != set.TotalCards {
			t.Errorf("Case %d: Expected %v, but got %v", index, test.totalCards, set.TotalCards)
		}
		if test.standardLegal != set.StandardLegal {
			t.Errorf("Case %d: Expected %v, but got %v", index, test.standardLegal, set.StandardLegal)
		}
		if test.expandedLegal != set.ExpandedLegal {
			t.Errorf("Case %d: Expected %v, but got %v", index, test.expandedLegal, set.ExpandedLegal)
		}
		if !strings.EqualFold(test.symbolURL, set.SymbolURL) {
			t.Errorf("Case %d: Expected %v, but got %v", index, test.symbolURL, set.SymbolURL)
		}
		if !strings.EqualFold(test.logoURL, set.LogoURL) {
			t.Errorf("Case %d: Expected %v, but got %v", index, test.logoURL, set.LogoURL)
		}

		if !strings.EqualFold(test.releasedDate, string(set.ReleasedDate)) {
			t.Errorf("Case %d: Expected %v, but got %v", index, test.releasedDate, string(set.ReleasedDate))
		}
		if !strings.EqualFold(test.updatedAt, string(set.UpdatedAt)) {
			t.Errorf("Case %d: Expected %v, but got %v", index, test.updatedAt, string(set.UpdatedAt))
		}

	}
}
