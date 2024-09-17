package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/jimsam/pokedexcli/pokecache"
)

type MapResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Resource string
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (r MapResponse) GetResource(resourceURL string, cache *pokecache.Cache) (interface{}, error) {
	url, err := url.Parse(resourceURL)
	if err != nil {
		return r, fmt.Errorf("The was an error with parsing %w", err)
	}
	query := url.Query()
	if len(query) == 0 {
		query.Add("limit", "0")
		query.Add("offset", "0")
		url.RawQuery = query.Encode()
	}

	data, found := cache.Get(url.String())
	if !found {
		fmt.Println("Fetching from web...")
		data, err = fetchFromWeb(url.String(), cache)
	}
	err = json.Unmarshal(data, &r)
	if err != nil {
		return r, err
	}
	r.Resource = "locations"
	printData(r)
	return r, nil
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

func printData(r MapResponse) error {
	for _, record := range r.Results {
		fmt.Println(record.Name)
	}
	return nil
}
