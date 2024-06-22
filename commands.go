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
