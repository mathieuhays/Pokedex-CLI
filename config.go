package main

import (
	"github.com/mathieuhays/pokedex-cli/internal/pokeapi"
)

type config struct {
	api                 *pokeapi.Client
	nextLocationURL     *string
	previousLocationURL *string
}
