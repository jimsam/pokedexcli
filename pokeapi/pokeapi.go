package pokeapi

import (
	"errors"
)

const baseURL = "https://pokeapi.co/api/v2/"

var resources = map[string]string{
	"map": baseURL + "location",
}

type FetchData interface {
	GetResource(resource string) (interface{}, error)
}

func ProcessRequest(r FetchData, resource string, lastResponse any) (interface{}, error) {
	resourceURL, err := getProperUrl(resource, lastResponse)
	if err != nil {
		return r, err
	}
	res, err := r.GetResource(resourceURL)
	if err != nil {
		return r, err
	}
	return res, nil
}

func getProperUrl(resource string, lastResponse any) (string, error) {
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
	return resourceURL, nil
}
