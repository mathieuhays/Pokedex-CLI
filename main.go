package main

import (
	"fmt"
	"github.com/mathieuhays/pokedex-cli/internal/pokeapi"
)

func main() {
	fmt.Println("  ===  P O K E D E X  ===  ")

	consoleRepl := newConsoleRepl(getCommands(), pokeapi.NewPokeAPI())
	consoleRepl.Listen()
}
