package main

import (
	"errors"
	"fmt"

	"github.com/jimsam/pokedexcli/pokeapi"
)

type commands struct {
	name        string
	description string
	callback    func(lastResponse *any, args []string) error
}

func getCommands() map[string]commands {
	return map[string]commands{
		"help": {
			name:        "help",
			description: "Displays available commands",
			callback: func(lastResponse *any, args []string) error {
				fmt.Println(PrintHelp())
				return nil
			},
		},
		"exit": {
			name:        "exit",
			description: "Exits pokedex.",
			callback:    func(lastResponse *any, args []string) error { return errors.New("exit") },
		},
		"map": {
			name:        "map",
			description: "Fetches the list of pokemon maps.",
			callback: func(lastResponse *any, args []string) error {
				r := pokeapi.MapResponse{}
				return pokeapi.ProcessRequest(r, "map", lastResponse, []string{})
			},
		},
		"mapb": {
			name:        "mapb",
			description: "Fetches the previous list of pokemon maps.",
			callback: func(lastResponse *any, args []string) error {
				r := pokeapi.MapResponse{}
				return pokeapi.ProcessRequest(r, "mapb", lastResponse, []string{})
			},
		},
		"visit": {
			name:        "visit",
			description: "Get the location areas of a map.",
			callback: func(lastResponse *any, args []string) error {
				r := pokeapi.LocationAreasResponse{}
				return pokeapi.ProcessRequest(r, "visit", lastResponse, args)
			},
		},
		"explore": {
			name:        "explore",
			description: "Explores a location for polemons.",
			callback: func(lastResponse *any, args []string) error {
				r := pokeapi.LocationAreaResponse{}
				return pokeapi.ProcessRequest(r, "explore", lastResponse, args)
			},
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon.",
			callback: func(lastResponse *any, args []string) error {
				r := pokeapi.PokemonResponse{}
				return pokeapi.ProcessRequest(r, "catch", lastResponse, args)
			},
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a caught pokemon.",
			callback: func(lastResponse *any, args []string) error {
				r := pokeapi.PokemonResponse{}
				return pokeapi.ProcessRequest(r, "inspect", lastResponse, args)
			},
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all caught pokemons.",
			callback: func(lastResponse *any, args []string) error {
				r := pokeapi.PokedexResponse{}
				return pokeapi.ProcessRequest(r, "pokedex", lastResponse, args)
			},
		},
	}
}
