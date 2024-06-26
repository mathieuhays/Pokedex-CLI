package pokeapi

func (c *Client) getLocationResourceURL() string {
	return baseUrl + "/location-area"
}

func (c *Client) getLocationAreaURL(area string) string {
	return baseUrl + "/location-area/" + area
}

func (c *Client) listLocationsWithURL(url string) (object NamedApiResourceList, err error) {
	err = c.requestWithCache(url, &object)
	if err == nil {
		object.Client = c
	}
	return
}

func (c *Client) ListLocations() (object NamedApiResourceList, err error) {
	return c.listLocationsWithURL(c.getLocationResourceURL())
}

func (c *Client) GetLocationAreaDetails(area string) (object LocationArea, err error) {
	err = c.requestWithCache(c.getLocationAreaURL(area), &object)
	return
}
