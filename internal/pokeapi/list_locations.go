package pokeapi

import (
	"encoding/json"
	"github.com/mathieuhays/pokedex-cli/internal/pokeapi/structs"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (structs.Resource, error) {
	url := baseUrl + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return structs.Resource{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return structs.Resource{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return structs.Resource{}, err
	}

	locations := structs.Resource{}
	err = json.Unmarshal(dat, &locations)
	if err != nil {
		return structs.Resource{}, err
	}

	return locations, nil
}
