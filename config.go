package main

import (
	"github.com/mathieuhays/pokedex-cli/internal/pokeapi"
	"github.com/mathieuhays/pokedex-cli/internal/pokedex"
)

type config struct {
	api          *pokeapi.Client
	pokedex      *pokedex.Pokedex
	lastLocation *pokeapi.NamedApiResourceList
}
