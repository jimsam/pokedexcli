# Pokedexcli

A GO application, to interact with the Pokemon API at `https://pokeapi.co/` from your CLI.

## How to use it.
If you have GO installed you can download the code and use `go run .` or build it with `go build .` and then just run the file `./pokedexcli`

### Available commands
* help -> will display the available commands
* map -> first 20 locations, any consecutive map command will return the next 20 results
* mapb -> (map back) will return the previous 20 locations
* exit -> exits

### Next steps
* Add caching so no calls to API will occur if the information is already fetched
* Caching will invalidate after 5 minutes.
