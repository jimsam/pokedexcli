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
	"explore": baseURL + "location-area/",
}

type FetchData interface {
	GetResource(resourceURL string, cache *pokecache.Cache) (interface{}, error)
}

var cache *pokecache.Cache

func init() {
	cache = pokecache.NewCache(time.Minute * 5)
}

func ProcessRequest(r FetchData, resource string, lastResponse *any, args []string) error {
	resourceURL, err := getProperUrl(resource, *lastResponse, args)
	if err != nil {
		return err
	}

	fmt.Println(resourceURL)
	*lastResponse, err = r.GetResource(resourceURL, cache)
	if err != nil {
		return err
	}
	return nil
}

func getProperUrl(resource string, lastResponse any, args []string) (string, error) {
	if lastResponse != nil {
		switch v := lastResponse.(type) {
		case MapResponse:
			if v.Resource == "locations" && resource == "map" {
				return v.Next, nil
			} else if v.Resource == "locations" && resource == "mapb" {
				url, ok := v.Previous.(string)
				if ok {
					return url, nil
				}
				return "", errors.New("You are at the first page!")
			}
		}
	}

	resourceURL, found := resources[resource]
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

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return []byte{}, err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		cache.Add(resourceURL, body)
		return body, nil
	} else {
		return []byte{}, fmt.Errorf("The was a response with %d code with message: %v", res.StatusCode, res.Body)
	}
}
