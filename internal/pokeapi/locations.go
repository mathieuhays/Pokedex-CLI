package pokeapi

import (
	"github.com/mathieuhays/pokedex-cli/internal/pokeapi/structs"
)

func (c *Client) GetLocationResourceURL() string {
	return baseUrl + "/location-area"
}

func (c *Client) GetLocationAreaURL(area string) string {
	return baseUrl + "/location-area/" + area
}

func (c *Client) ListLocations(url string) (structs.NamedApiResourceList, error) {
	var object structs.NamedApiResourceList
	err := c.requestWithCache(url, &object)
	if err != nil {
		return structs.NamedApiResourceList{}, err
	}
	return object, nil
}

func (c *Client) GetLocationAreaDetails(url string) (structs.LocationArea, error) {
	var object structs.LocationArea
	err := c.requestWithCache(url, &object)
	if err != nil {
		return structs.LocationArea{}, err
	}
	return object, nil
}
