package main

import (
	"fmt"
	"github.com/mathieuhays/pokedex-cli/internal/pokeapi"
)

func main() {
	fmt.Println("  ===  P O K E D E X  ===  ")

	cfg := &config{
		api: pokeapi.NewPokeAPI(),
	}

	consoleRepl := newConsoleRepl(getCommands(), cfg)
	consoleRepl.Listen()
}
