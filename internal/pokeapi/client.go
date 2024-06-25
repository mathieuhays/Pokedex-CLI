package pokeapi

import (
	"github.com/mathieuhays/pokedex-cli/internal/pokecache"
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
