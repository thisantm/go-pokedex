package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Locations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocations(pageURL *string) (Locations, error) {
	endpoint := "/location-area/"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return Locations{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return Locations{}, err
	}
	defer resp.Body.Close()

	if resp.Status >= "400" {
		return Locations{}, errors.New("error: Not Found")
	}

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return Locations{}, err
	}

	locations := Locations{}
	err = json.Unmarshal(bodyBytes, &locations)

	if err != nil {
		return Locations{}, err
	}

	return locations, nil
}
