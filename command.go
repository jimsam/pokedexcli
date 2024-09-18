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
			description: "Exits pokedex",
			callback:    func(lastResponse *any, args []string) error { return errors.New("exit") },
		},
		"map": {
			name:        "map",
			description: "Fetches the list of pokemon maps",
			callback: func(lastResponse *any, args []string) error {
				r := pokeapi.MapResponse{}
				return pokeapi.ProcessRequest(r, "map", lastResponse, []string{})
			},
		},
		"mapb": {
			name:        "mapb",
			description: "Fetches the previous list of pokemon maps",
			callback: func(lastResponse *any, args []string) error {
				r := pokeapi.MapResponse{}
				return pokeapi.ProcessRequest(r, "mapb", lastResponse, []string{})
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
	}
}
