package pokeapi

import (
	"fmt"

	"github.com/jimsam/pokedexcli/pokecache"
)

type PokedexResponse struct {
	Name string
}

func (r PokedexResponse) GetResource(resourceURL string, cache *pokecache.Cache, action string, pokedex map[string]pokecache.Pokedex, args []string) (interface{}, error) {
	for _, val := range pokedex {
		fmt.Println("- ", val.Name)
	}
	return r, nil
}
