package pokeapi

import (
	"encoding/json"
	"errors"
	"github.com/mathieuhays/pokedex-cli/internal/pokecache"
	"io"
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(timeout time.Duration, cache *pokecache.Cache) Client {
	return Client{
		httpClient: http.Client{Timeout: timeout},
		cache:      cache,
	}
}

func (c *Client) requestWithCache(url string, obj interface{}) error {
	dat, exists := c.cache.Get(url)
	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode == 404 {
			return errors.New("not found")
		} else if resp.StatusCode >= 300 {
			return errors.New("unexpected error")
		}

		dat, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		c.cache.Add(url, dat)
	}

	return json.Unmarshal(dat, &obj)
}
