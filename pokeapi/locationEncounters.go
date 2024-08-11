package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Encounters struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetLocationEncounters(locationName string) (Encounters, error) {
	endpoint := "/location-area/"
	fullURL := baseURL + endpoint + locationName

	bodyBytes, ok := c.cache.Get(fullURL)

	if ok {
		encounters := Encounters{}
		err := json.Unmarshal(bodyBytes, &encounters)

		if err != nil {
			return Encounters{}, err
		}
		return encounters, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return Encounters{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return Encounters{}, err
	}
	defer resp.Body.Close()

	if resp.Status >= "400" {
		return Encounters{}, errors.New("error: Not Found")
	}

	bodyBytes, err = io.ReadAll(resp.Body)

	if err != nil {
		return Encounters{}, err
	}

	encounters := Encounters{}
	err = json.Unmarshal(bodyBytes, &encounters)

	if err != nil {
		return Encounters{}, err
	}

	go c.cache.Add(fullURL, bodyBytes)

	return encounters, nil
}
