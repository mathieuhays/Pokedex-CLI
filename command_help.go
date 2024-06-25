package main

import "fmt"

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
