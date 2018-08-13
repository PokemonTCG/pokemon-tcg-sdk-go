package pokemontcgsdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	//EndPointTypes used to obtain list of pokemon types from the pokemontcg api
	EndPointTypes = EndPoint + "/types/"

	//EndPointSubTypes used to obtain list of pokemon card subtypes
	EndPointSubTypes = EndPoint + "/subtypes/"

	//EndPointSuperTypes used to obtain list of pokemon case supertypes
	EndPointSuperTypes = EndPoint + "/supertypes/"
)

func getTypes(endPoint string, mapKey string) (types []string, err error) {
	typesMap := make(map[string][]string)
	resp, err := http.Get(endPoint)
	if err != nil {
		return types, fmt.Errorf("unable to http.Get types: %s", err.Error())
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&typesMap)
	if err != nil {
		return types, fmt.Errorf("Unable to decode json response: %s", err.Error())
	}

	var ok bool
	types, ok = typesMap[mapKey]
	if !ok {
		return types, fmt.Errorf("unable to get types from map: %s", err.Error())
	}
	return types, err

}

// GetTypes returns a list of pokemon types.
func GetTypes() ([]string, error) {
	return getTypes(EndPointTypes, "types")
}

// GetSubTypes returns a list of pokemon card subtypes
func GetSubTypes() ([]string, error) {
	return getTypes(EndPointSubTypes, "subtypes")
}

// GetSuperTypes returns a list of pokemon card supertypes
func GetSuperTypes() ([]string, error) {
	return getTypes(EndPointSuperTypes, "supertypes")
}
