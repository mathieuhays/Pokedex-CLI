package main

type cliCommand struct {
	name        string
	description string
	callback    func(r *repl, args []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "List next locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous locations",
			callback:    commandPreviousMap,
		},
		"explore": {
			name:        "explore",
			description: "Explore an area and see which pokemon can be found there. Use the map command to get a list of areas.",
			callback:    commandExplore,
		},
		"pokemon": {
			name:        "pokemon",
			description: "gets stats about a pokemon",
			callback:    commandPokemon,
		},
	}
}
