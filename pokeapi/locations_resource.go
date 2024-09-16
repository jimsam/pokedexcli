package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
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

func (r MapResponse) GetResource(resourceURL string) (interface{}, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", resourceURL, nil)
	if err != nil {
		return r, err
	}
	res, err := client.Do(req)
	if err != nil {
		return r, err
	}

	if res.StatusCode < 400 {
		body, err := io.ReadAll(res.Body)
		defer res.Body.Close()
		if err != nil {
			return r, err
		}

		err = json.Unmarshal(body, &r)
		if err != nil {
			return r, err
		}
		r.Resource = "locations"
		printData(r)
		return r, nil
	} else {
		return r, errors.New("An error occured")
	}
}

func printData(r MapResponse) error {
	for _, record := range r.Results {
		fmt.Println(record.Name)
	}
	return nil
}
