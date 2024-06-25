package pokeapi

import (
	"encoding/json"
	"github.com/mathieuhays/pokedex-cli/internal/pokeapi/structs"
	"io"
	"net/http"
)

func (c *Client) GetLocationURL() string {
	return baseUrl + "/location-area"
}

func (c *Client) ListLocations(url string) (structs.Resource, error) {
	dat, exists := c.cache.Get(url)
	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return structs.Resource{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return structs.Resource{}, err
		}
		defer resp.Body.Close()

		dat, err = io.ReadAll(resp.Body)
		if err != nil {
			return structs.Resource{}, err
		}

		c.cache.Add(url, dat)
	}

	locations := structs.Resource{}
	err := json.Unmarshal(dat, &locations)
	if err != nil {
		return structs.Resource{}, err
	}

	return locations, nil
}
