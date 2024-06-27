package pokedex

import (
	"errors"
	"github.com/mathieuhays/pokedex-cli/internal/pokeapi"
)

type Pokedex struct {
	store    map[string]pokeapi.Pokemon
	attempts map[string]int
}

func (p *Pokedex) Has(name string) bool {
	_, ok := p.store[name]
	return ok
}

func (p *Pokedex) Add(pokemon pokeapi.Pokemon) error {
	if p.Has(pokemon.Name) {
		return errors.New("pokemon already in pokedex")
	}

	p.store[pokemon.Name] = pokemon

	return nil
}

func (p *Pokedex) Get(name string) (pokemon pokeapi.Pokemon, exists bool) {
	if !p.Has(name) {
		return
	}

	return p.store[name], true
}

func (p *Pokedex) GetAll() map[string]pokeapi.Pokemon {
	return p.store
}

func (p *Pokedex) RegisterAttempt(name string) {
	if _, ok := p.attempts[name]; !ok {
		p.attempts[name] = 0
	}

	p.attempts[name]++
}

func (p *Pokedex) GetNumberOfAttempts(name string) int {
	if attempts, ok := p.attempts[name]; ok {
		return attempts
	}

	return 1
}

func NewPokedex() *Pokedex {
	return &Pokedex{
		store:    make(map[string]pokeapi.Pokemon),
		attempts: make(map[string]int),
	}
}
