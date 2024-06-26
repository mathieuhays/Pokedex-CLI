package main

import (
	"github.com/mathieuhays/pokedex-cli/internal/pokeapi"
)

type config struct {
	api          *pokeapi.Client
	lastLocation *pokeapi.NamedApiResourceList
}
