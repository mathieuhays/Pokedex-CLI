package main

import (
	"fmt"
	"github.com/mathieuhays/pokedex-cli/internal/pokeapi"
	"github.com/mathieuhays/pokedex-cli/internal/pokecache"
	pokedex2 "github.com/mathieuhays/pokedex-cli/internal/pokedex"
	"time"
)

func main() {
	fmt.Println("  ===  P O K E D E X  ===  ")

	cache := pokecache.NewCache(time.Minute*15, time.Minute*15)
	pokeClient := pokeapi.NewClient(time.Second*2, cache)
	pokedex := pokedex2.NewPokedex()
	cfg := &config{
		api:     &pokeClient,
		pokedex: pokedex,
	}

	consoleRepl := newConsoleRepl(getCommands(), cfg)
	consoleRepl.Listen()
}
