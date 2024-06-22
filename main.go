package main

import (
	"fmt"
)

func main() {
	fmt.Println("  ===  P O K E D E X  ===  ")

	consoleRepl := newConsoleRepl(getCommands())
	consoleRepl.Listen()
}
