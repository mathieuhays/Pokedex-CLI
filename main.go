package main

import (
	"fmt"
	"github.com/mathieuhays/pokedex-cli/internal/pokeapi"
	"time"
)

func main() {
	fmt.Println("  ===  P O K E D E X  ===  ")

	pokeClient := pokeapi.NewClient(time.Second * 2)
	cfg := &config{
		api: &pokeClient,
	}

	consoleRepl := newConsoleRepl(getCommands(), cfg)
	consoleRepl.Listen()
}
