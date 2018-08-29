package pokemontcgsdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	//EndPointSets used to obtain information about card sets from the pokemontcg api
	EndPointSets = EndPoint + "/sets"
	//EndPointSetID used to obtain a set by ID
	EndPointSetID = EndPoint + "/sets/"
)

// Set describes the official pokemon sets of cards
type Set struct {
	Code          string          `json:"code"`
	PtcgoCode     string          `json:"ptcgoCode"`
	Name          string          `json:"name"`
	Series        string          `json:"series"`
	TotalCards    int             `json:"totalCards"`
	StandardLegal bool            `json:"standardLegal"`
	ExpandedLegal bool            `json:"expandedLegal"`
	SymbolURL     string          `json:"symbolUrl"`
	LogoURL       string          `json:"logoUrl"`
	ReleasedDate  json.RawMessage `json:"releasedDate"`
	UpdatedAt     json.RawMessage `json:"updatedAt"`
}

// GetSets allows you to search and filter for sets using the parameter
// list in the pokemontcg.io docs website (https://docs.pokemontcg.io/#api_v1sets_list)
func GetSets(params map[string]string) (sets []Set, err error) {
	urlQuery, err := formatQuery(EndPointSets, params)
	if err != nil {
		return sets, err
	}

	setsMap := make(map[string][]Set)
	resp, err := http.Get(urlQuery)
	if err != nil {
		return sets, fmt.Errorf("unable to http.Get cards: %s", err.Error())
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&setsMap)
	if err != nil {
		return sets, fmt.Errorf("Unable to decode json response: %s", err.Error())
	}

	var ok bool
	sets, ok = setsMap["sets"]
	if !ok {
		return sets, fmt.Errorf("unable to get cards from cards map: %s", err.Error())
	}
	return sets, err
}

// GetSetByID returns as single pokemon card.
func GetSetByID(code string) (set Set, err error) {
	sets := make(map[string]Set)
	resp, err := http.Get(EndPointSetID + code)
	if err != nil {
		return set, fmt.Errorf("unable to http.Get card: %s", err.Error())
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&sets)
	if err != nil {
		return set, fmt.Errorf("Unable to decode json response: %s", err.Error())
	}

	var ok bool
	set, ok = sets["set"]
	if !ok {
		return set, fmt.Errorf("unable to get set from sets map: %s", err.Error())
	}
	return set, err

}
