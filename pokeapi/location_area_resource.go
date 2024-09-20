package pokeapi

import (
	"encoding/json"
	"fmt"

	"github.com/jimsam/pokedexcli/pokecache"
)

type LocationAreaResponse struct {
	Resource             string
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (r LocationAreaResponse) GetResource(resourceURL string, cache *pokecache.Cache) (interface{}, error) {
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
	r.Resource = "location-area"
	printPokemonInArea(r)
	return r, nil
}

func printPokemonInArea(r LocationAreaResponse) {
	for _, val := range r.PokemonEncounters {
		fmt.Println("- ", val.Pokemon.Name)
	}
}
