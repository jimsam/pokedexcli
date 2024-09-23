# Pokedexcli

A Go application, to interact with the Pokemon API at `https://pokeapi.co/` from your CLI.

## How to use it.
If you have Go installed you can download the code and use `go run .` or build it with `go build .` and then just run the file `./pokedexcli`

### Available commands
* help -> Lists the available commands.
* map -> Lists the first 20 locations, any consecutive map command will return the next 20 results
* mapb -> Lists the previous 20 locations, if not in first page.
* visit -> Lists all areas available in that location.
* explore -> Lists all pokemons found in that area. 
* catch -> Tries to catch a pokemon.
* inspect <pokemon name> -> If have caught that pokemon it returns details about it.
* pokedex -> Lists all caught pokemons.
* exit -> exits

### WIP
