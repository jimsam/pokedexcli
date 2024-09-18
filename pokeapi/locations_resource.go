package pokeapi

import (
	"encoding/json"
	"fmt"

	"github.com/jimsam/pokedexcli/pokecache"
)

type MapResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Resource string
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (r MapResponse) GetResource(resourceURL string, cache *pokecache.Cache) (interface{}, error) {
	data, found := cache.Get(resourceURL)
	var err error
	if !found {
		fmt.Println("Fetching from web...")
		data, err = fetchFromWeb(resourceURL, cache)
	}
	err = json.Unmarshal(data, &r)
	if err != nil {
		return r, err
	}
	r.Resource = "locations"
	printLocations(r)
	return r, nil
}

func printLocations(r MapResponse) error {
	for _, record := range r.Results {
		fmt.Println(record.Name)
	}
	return nil
}
