package pokemontcgsdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	//EndPointSets used to obtain information about card sets from the pokemontcg api
	EndPointSets = EndPoint + "/sets/"
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

// GetSetByID returns as single pokemon card.
func GetSetByID(code string) (set Set, err error) {
	sets := make(map[string]Set)
	resp, err := http.Get(EndPointSets + code)
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
