package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type repl struct {
	scanner  *bufio.Scanner
	out      io.Writer
	commands map[string]cliCommand
}

const promptPrefix = "pokedex > "

func (r *repl) Listen() {
	for {
		_, _ = fmt.Fprintf(r.out, "\n%s", promptPrefix)

		if !r.scanner.Scan() {
			break
		}

		words := cleanInput(r.scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := r.commands[commandName]
		if exists {
			err := command.callback(r.out, words[1:])
			if err != nil {
				_, _ = fmt.Fprintf(r.out, "%s error: %s", command.name, err.Error())
			}
		} else {
			_, _ = fmt.Fprintln(r.out, "unknown command")
		}
	}
}

func newConsoleRepl(commands map[string]cliCommand) *repl {
	return &repl{
		scanner:  bufio.NewScanner(os.Stdin),
		out:      os.Stdout,
		commands: commands,
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
