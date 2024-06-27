package pokeapi

func (c *Client) GetPokemon(name string) (pokemon Pokemon, err error) {
	url := baseUrl + "/pokemon/" + name
	err = c.requestWithCache(url, &pokemon)
	return
}
