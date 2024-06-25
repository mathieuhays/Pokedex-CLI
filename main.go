package main

import (
	"fmt"
	"github.com/mathieuhays/pokedex-cli/internal/pokeapi"
	"github.com/mathieuhays/pokedex-cli/internal/pokecache"
	"time"
)

func main() {
	fmt.Println("  ===  P O K E D E X  ===  ")

	cache := pokecache.NewCache(time.Minute * 15)
	pokeClient := pokeapi.NewClient(time.Second*2, cache)
	cfg := &config{
		api: &pokeClient,
	}

	consoleRepl := newConsoleRepl(getCommands(), cfg)
	consoleRepl.Listen()
}
