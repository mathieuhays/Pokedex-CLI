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
	}
}
