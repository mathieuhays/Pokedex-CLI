package main

import (
	"fmt"
	"os"
)

func commandExit(r *repl, args []string) error {
	_, _ = fmt.Fprintln(r.out, "Goodbye")
	os.Exit(0)
	return nil
}
