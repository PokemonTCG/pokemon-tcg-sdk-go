package pokemontcgsdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	//EndPointTypes used to obtain infromation about pokemone types from the pokemontcg api
	EndPointTypes = EndPoint + "/types/"
)

// GetTypes returns a list of pokemon types.
func GetTypes() (types []string, err error) {
	typesMap := make(map[string][]string)
	resp, err := http.Get(EndPointTypes)
	if err != nil {
		return types, fmt.Errorf("unable to http.Get types: %s", err.Error())
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&typesMap)
	if err != nil {
		return types, fmt.Errorf("Unable to decode json response: %s", err.Error())
	}

	var ok bool
	types, ok = typesMap["types"]
	if !ok {
		return types, fmt.Errorf("unable to get types from typesMap: %s", err.Error())
	}
	return types, err

}
