package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	fmt.Println("  ===  P O K E D E X  ===  ")
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	for {
		fmt.Print("\npokedex > ")
		scanner.Scan()
		input := scanner.Text()

		if _, ok := commands[input]; !ok {
			fmt.Println("command not found")
			_ = commandHelp()
			continue
		}

		err := commands[input].callback()
		if err != nil {
			fmt.Printf("%s error: %s", commands[input].name, err.Error())
		}
	}
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

func commandHelp() error {
	commands := getCommands()
	fmt.Println("Usage:")

	if len(commands) == 0 {
		fmt.Println("No commands found")
	}

	for input, command := range commands {
		fmt.Printf("%s: %s\n", input, command.description)
	}

	return nil
}

func commandExit() error {
	fmt.Println("Lemme exit")
	os.Exit(0)
	return nil
}
