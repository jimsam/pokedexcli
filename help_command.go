package main

func PrintHelp() string {
	return `
Welcome to the Pokedex!
Usage:

map: Lists all location that you can visit in the game (20 location per page)
mapb: If you previous command was map and not in the first page it takes you back one page
visit: Lists all area available to explore in that location
explore: Lists all pokemon found in that location
catch: Try to catch a pokemon
inspect <pokemon name>: If you have caught that pokemon then you get informations of that pokemon
pokedex: A list of all the pokemons that you have caught
help: Displays a help message
exit: Exit the Pokedex
`
}
