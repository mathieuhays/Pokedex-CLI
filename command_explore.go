package main

import (
	"errors"
	"fmt"
)

func commandExplore(r *repl, args []string) error {
	if len(args) == 0 {
		return errors.New("please specify a location to explore")
	}
	if len(args) > 1 {
		return errors.New("too many arguments")
	}

	name := args[0]
	details, err := r.config.api.GetLocationAreaDetails(name)
	if err != nil {
		return err
	}

	_, _ = fmt.Fprintf(r.out, "Exploring %v...\n", details.Name)

	if len(details.PokemonEncounters) == 0 {
		_, _ = fmt.Fprintln(r.out, "No pokemon found in this area")
		return nil
	}

	_, _ = fmt.Fprintln(r.out, "Found Pokemon:")
	for _, pokemon := range details.PokemonEncounters {
		_, _ = fmt.Fprintf(r.out, "- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
