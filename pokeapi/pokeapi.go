package pokeapi

import (
	"net/http"
	"time"

	"github.com/thisantm/go-pokedex/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(cacheCleanupTime time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheCleanupTime),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
