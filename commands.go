package main

import (
	"fmt"
	"os"
)

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
		// @TODO remove debug commands eventually?
		"debug": {
			name:        "debug",
			description: "Debug api information",
			callback:    commandDebug,
		},
		"debug_cache": {
			name:        "debug_cache",
			description: "Debug api cache in details",
			callback:    commandDebugCache,
		},
	}
}

func commandHelp(r *repl, args []string) error {
	commands := r.commands
	_, _ = fmt.Fprintln(r.out, "Usage:")

	if len(commands) == 0 {
		_, _ = fmt.Fprintln(r.out, "No commands found")
	}

	for _, command := range commands {
		_, _ = fmt.Fprintf(r.out, "%s: %s\n", command.name, command.description)
	}

	return nil
}

func commandExit(r *repl, args []string) error {
	_, _ = fmt.Fprintln(r.out, "Goodbye")
	os.Exit(0)
	return nil
}

func commandMap(r *repl, args []string) error {
	resource, err := r.api.GetLocations()
	if err != nil {
		return err
	}

	for _, result := range resource.Results {
		_, _ = fmt.Fprintf(r.out, "%s\n", result.Name)
	}

	return nil
}

func commandPreviousMap(r *repl, args []string) error {
	resource, err := r.api.GetPreviousLocations()
	if err != nil {
		return err
	}

	for _, result := range resource.Results {
		_, _ = fmt.Fprintf(r.out, "%s\n", result.Name)
	}

	return nil
}

func commandDebug(r *repl, args []string) error {
	r.api.Debug(r.out)
	return nil
}

func commandDebugCache(r *repl, args []string) error {
	r.api.DebugCache(r.out)
	return nil
}
