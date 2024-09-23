package pokeapi

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/jimsam/pokedexcli/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2/"

var resources = map[string]string{
	"map":     baseURL + "location",
	"visit":   baseURL + "location/",
	"explore": baseURL + "location-area/",
	"catch":   baseURL + "pokemon/",
	"inspect": baseURL + "pokemon/",
	"species": baseURL + "pokemon-species/",
	"pokedex": "",
}

type FetchData interface {
	GetResource(resourceURL string, cache *pokecache.Cache, action string, pokedex map[string]pokecache.Pokedex, args []string) (interface{}, error)
}

var cache *pokecache.Cache
var pokedex map[string]pokecache.Pokedex

func init() {
	cache = pokecache.NewCache(time.Minute * 5)
	pokedex = pokecache.NewPokedex()
}

func ProcessRequest(r FetchData, action string, lastResponse *any, args []string) error {
	resourceURL, err := getProperUrl(action, *lastResponse, args)
	if err != nil {
		return err
	}

	*lastResponse, err = r.GetResource(resourceURL, cache, action, pokedex, args)
	if err != nil {
		return err
	}
	return nil
}

func getProperUrl(action string, lastResponse any, args []string) (string, error) {
	if lastResponse != nil {
		switch v := lastResponse.(type) {
		case MapResponse:
			if v.Resource == "locations" && action == "map" {
				return v.Next, nil
			} else if v.Resource == "locations" && action == "mapb" {
				url, ok := v.Previous.(string)
				if ok {
					return url, nil
				}
				return "", errors.New("You are at the first page!")
			}
		}
	}

	resourceURL, found := resources[action]
	if !found {
		return "", errors.New("Requested resource was not found!")
	}
	if len(args) > 1 {
		resourceURL += args[1]
		return resourceURL, nil
	}
	resourceURL = addFirstPagePagination(resourceURL)
	return resourceURL, nil
}

func addFirstPagePagination(url string) string {
	return url + "?offset=0&limit=20"
}

func fetchFromWeb(resourceURL string, cache *pokecache.Cache) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", resourceURL, nil)
	if err != nil {
		return []byte{}, err
	}

	res, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return []byte{}, err
		}
		cache.Add(resourceURL, body)
		return body, nil
	} else {
		return []byte{}, fmt.Errorf("The was a response with %d code", res.StatusCode)
	}
}
