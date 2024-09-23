package pokeapi

import (
	"encoding/json"
	"fmt"

	"github.com/jimsam/pokedexcli/pokecache"
)

type LocationAreasResponse struct {
	Resource string
	Areas    []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"areas"`
	GameIndices []struct {
		GameIndex  int `json:"game_index"`
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
	} `json:"game_indices"`
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	Region struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"region"`
}

func (r LocationAreasResponse) GetResource(resourceURL string, cache *pokecache.Cache, action string, pokedex map[string]pokecache.Pokedex, args []string) (interface{}, error) {
	data, found := cache.Get(resourceURL)
	var err error
	if !found {
		fmt.Println("Fetching from web...")
		data, err = fetchFromWeb(resourceURL, cache)
	}
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &r)
	if err != nil {
		return r, err
	}

	r.Resource = "location-areas"
	printLocationAreas(r)
	return r, nil
}

func printLocationAreas(r LocationAreasResponse) {
	for _, val := range r.Areas {
		fmt.Println("- ", val.Name)
	}
}
