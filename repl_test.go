package main

import (
	"bufio"
	"bytes"
	"fmt"
	"slices"
	"testing"
)

// This feels more like an integration test so maybe this should be refactored
func TestRepl_Listen(t *testing.T) {
	t.Run("empty prompt", func(t *testing.T) {
		bufIn := bytes.Buffer{}
		bufOut := bytes.Buffer{}
		testRepl := &repl{
			scanner: bufio.NewScanner(&bufIn),
			out:     &bufOut,
			commands: map[string]cliCommand{
				"help": {
					name:        "help",
					description: "help",
					callback: func(r *repl, args []string) error {
						return nil
					},
				},
			},
		}

		bufIn.WriteString("\n")
		testRepl.Listen()
		expected := fmt.Sprintf("\n%s\n%s", promptPrefix, promptPrefix)
		actual := bufOut.String()

		if expected != actual {
			t.Errorf("got %q want %q", actual, expected)
		}
	})

	t.Run("prompt with command", func(t *testing.T) {
		bufIn := bytes.Buffer{}
		bufOut := bytes.Buffer{}
		testRepl := &repl{
			scanner: bufio.NewScanner(&bufIn),
			out:     &bufOut,
			commands: map[string]cliCommand{
				"command": {
					name:        "command",
					description: "command",
					callback: func(r *repl, args []string) error {
						_, _ = fmt.Fprintln(r.out, "test")
						return nil
					},
				},
			},
		}

		bufIn.WriteString("command")
		testRepl.Listen()
		expected := fmt.Sprintf("\n%stest\n\n%s", promptPrefix, promptPrefix)
		actual := bufOut.String()

		if expected != actual {
			t.Errorf("got %q want %q", actual, expected)
		}
	})

	t.Run("prompt with unknown command", func(t *testing.T) {
		bufIn := bytes.Buffer{}
		bufOut := bytes.Buffer{}
		testRepl := &repl{
			scanner: bufio.NewScanner(&bufIn),
			out:     &bufOut,
			commands: map[string]cliCommand{
				"command": {
					name:        "command",
					description: "command",
					callback: func(r *repl, args []string) error {
						_, _ = fmt.Fprintln(r.out, "test")
						return nil
					},
				},
			},
		}

		bufIn.WriteString("random")
		testRepl.Listen()
		expected := fmt.Sprintf("\n%sunknown command\n\n%s", promptPrefix, promptPrefix)
		actual := bufOut.String()

		if expected != actual {
			t.Errorf("got %q want %q", actual, expected)
		}
	})
}

func TestCleanInput(t *testing.T) {
	t.Run("normalize case", func(t *testing.T) {
		expected := []string{"help"}
		actual := cleanInput("Help")

		if !slices.Equal(expected, actual) {
			t.Errorf("got %v want %v", actual, expected)
		}
	})

	t.Run("multiple words", func(t *testing.T) {
		expected := []string{"help", "test"}
		actual := cleanInput("Help test")

		if !slices.Equal(expected, actual) {
			t.Errorf("got %v want %v", actual, expected)
		}
	})

	t.Run("weird spacing", func(t *testing.T) {
		expected := []string{"help", "test"}
		actual := cleanInput(" Help    test ")

		if !slices.Equal(expected, actual) {
			t.Errorf("got %v want %v", actual, expected)
		}
	})
}
