package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"

	"github.com/jimsam/pokedexcli/pokecache"
)

func (r PokemonResponse) GetResource(resourceURL string, cache *pokecache.Cache, action string, pokedex map[string]pokecache.Pokedex, args []string) (interface{}, error) {
	if action == "catch" {
		return r.catchPokemon(resourceURL, cache, pokedex)
	} else {
		data, found := pokedex[args[1]]
		if !found {
			return nil, fmt.Errorf("you have not caught that pokemon")
		}
		printPokemon(data)
		return r, nil
	}
}

func (r PokemonResponse) catchPokemon(resourceURL string, cache *pokecache.Cache, pokedex map[string]pokecache.Pokedex) (interface{}, error) {
	data, found := cache.Get(resourceURL)
	var err error
	if !found {
		data, err = fetchFromWeb(resourceURL, cache)
		if err != nil {
			return nil, err
		}
	}

	err = json.Unmarshal(data, &r)
	if err != nil {
		return nil, err
	}
	r.Resource = "pokemon"
	captureRate, err := getFoundPokemonCaptureRate(r.Species.Name, cache)
	if err != nil {
		return r, err
	}

	if tryToCatch(captureRate, r.Name) {
		pokedex[r.Name] = pokecache.Pokedex{
			Height:          r.Height,
			Weight:          r.Weight,
			Name:            r.Name,
			Hp:              findFromStats(r.Stats, "hp"),
			Attack:          findFromStats(r.Stats, "attack"),
			Defense:         findFromStats(r.Stats, "defense"),
			Special_attack:  findFromStats(r.Stats, "special-attack"),
			Special_defense: findFromStats(r.Stats, "special-defense"),
			Speed:           findFromStats(r.Stats, "speed"),
			Types:           getAllTypes(r.Types),
		}
	}

	return r, nil
}

func printPokemon(pokedexData pokecache.Pokedex) {
	fmt.Printf("%v\n", pokedexData)
	fmt.Println("Name: ", pokedexData.Name)
	fmt.Println("Height: ", pokedexData.Height)
	fmt.Println("Weight: ", pokedexData.Weight)
	fmt.Println("Stats:")
	fmt.Println("  -hp: ", pokedexData.Hp)
	fmt.Println("  -attack: ", pokedexData.Attack)
	fmt.Println("  -defense: ", pokedexData.Defense)
	fmt.Println("  -special-attack: ", pokedexData.Special_attack)
	fmt.Println("  -special-defense: ", pokedexData.Special_defense)
	fmt.Println("  -speed: ", pokedexData.Speed)
	fmt.Println("Types:")
	for _, tp := range pokedexData.Types {
		fmt.Println("  - ", tp)
	}
}

func getFoundPokemonCaptureRate(specie string, cache *pokecache.Cache) (int, error) {
	s := SpeciesResponse{}

	resourceURL, found := resources["species"]
	if !found {
		return 0, errors.New("Requested resource was not found!")
	}
	return s.GetCaptureRate(resourceURL+specie, cache)
}

func tryToCatch(captureRate int, pokemonName string) bool {
	fmt.Println("Throwing the pokeball to ", pokemonName, "...")
	catchScore := rand.Intn(256)

	if catchScore <= captureRate {
		fmt.Println(pokemonName, " was caught")
		return true
	} else {
		fmt.Println(pokemonName, " escaped!")
		return false
	}
}

func findFromStats(stats []Statistics, statName string) int {
	for _, statistic := range stats {
		if statistic.Stat.Name == statName {
			return statistic.BaseStat
		}
	}

	return -1
}

func getAllTypes(ptypes []PTypes) []string {
	var res []string
	for _, tp := range ptypes {
		res = append(res, tp.Type.Name)
	}
	return res
}
