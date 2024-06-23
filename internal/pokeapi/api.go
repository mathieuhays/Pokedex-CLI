package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mathieuhays/pokedex-cli/internal/pokeapi/structs"
	"io"
	"log"
	"net/http"
	"strings"
)

const baseUrl = "https://pokeapi.co/api/v2"

func request(endpoint string) (string, error) {
	res, err := http.Get(endpoint)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(res.Body)
	_ = res.Body.Close()

	if res.StatusCode > 299 {
		return "", errors.New(fmt.Sprintf("Response failed with status code: %v", res.StatusCode))
	}

	if err != nil {
		return "", err
	}

	return string(body), nil
}

type PokeAPI struct {
	cache                map[string]string
	lastLocationEndpoint *string
}

func NewPokeAPI() *PokeAPI {
	cache := make(map[string]string)
	return &PokeAPI{
		cache: cache,
	}
}

func (p *PokeAPI) normalizeURL(endpoint string) string {
	if strings.HasPrefix(endpoint, "http") {
		return endpoint
	}

	if !strings.HasPrefix(endpoint, "/") {
		endpoint = "/" + endpoint
	}

	return baseUrl + endpoint
}

func (p *PokeAPI) get(path string, obj interface{}) error {
	url := p.normalizeURL(path)

	if data, exists := p.cache[url]; exists {
		err := json.Unmarshal([]byte(data), &obj)
		if err == nil {
			return nil // if we get an error parsing JSON from the cache, let's try the request again
		}
	}

	body, err := request(url)
	if err != nil {
		return err
	}

	p.cache[url] = body
	return json.Unmarshal([]byte(body), &obj)
}

func (p *PokeAPI) getNextLocationURL() string {
	if p.lastLocationEndpoint != nil {
		var res structs.Resource
		err := p.get(*p.lastLocationEndpoint, &res)
		if err == nil {
			return res.Next
		}

		log.Printf("lastLocationEndpoint json error: %s", err.Error())
	}

	return "https://pokeapi.co/api/v2/location-area/"
}

func (p *PokeAPI) getPreviousLocationURL() (string, error) {
	if p.lastLocationEndpoint == nil {
		return "", errors.New("no previous location to load")
	}

	var res structs.Resource
	err := p.get(*p.lastLocationEndpoint, &res)
	if err != nil {
		return "", errors.New(fmt.Sprintf("could not process the last location. error: %s", err.Error()))
	}

	if res.Previous == nil {
		return "", errors.New("no previous location available")
	}

	return *res.Previous, nil
}

func (p *PokeAPI) GetLocations() (*structs.Resource, error) {
	url := p.getNextLocationURL()
	var res structs.Resource

	err := p.get(url, &res)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error retrieving locations: %s", err.Error()))
	}

	p.lastLocationEndpoint = &url

	return &res, nil
}

func (p *PokeAPI) GetPreviousLocations() (*structs.Resource, error) {
	url, err := p.getPreviousLocationURL()
	if err != nil {
		return nil, err
	}

	var res structs.Resource

	err = p.get(url, &res)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error retrieving locations: %s", err.Error()))
	}

	p.lastLocationEndpoint = &url

	return &res, nil
}

func (p *PokeAPI) Debug(out io.Writer) {
	if p.lastLocationEndpoint != nil {
		fmt.Fprintf(out, "has lastLocationEndpoint: %s\n", *p.lastLocationEndpoint)
	} else {
		fmt.Fprintln(out, "no lastLocationEndpoint set")
	}

	fmt.Fprintf(out, "has %d items in cache", len(p.cache))
}

func (p *PokeAPI) DebugCache(out io.Writer) {
	if len(p.cache) == 0 {
		fmt.Fprintln(out, "no items in cache")
		return
	}

	for key, val := range p.cache {
		fmt.Fprintf(out, "- key: %s value length: %d\n", key, len(val))
	}
}
