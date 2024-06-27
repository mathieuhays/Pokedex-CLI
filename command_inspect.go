package main

import (
	"errors"
	"fmt"
)

func commandInspect(r *repl, args []string) error {
	if len(args) == 0 {
		return errors.New("please specify the name of a pokemon")
	}
	if len(args) > 1 {
		return errors.New("too many arguments")
	}

	name := args[0]
	pokemon, exists := r.config.pokedex.Get(name)
	if !exists {
		_, _ = fmt.Fprintln(r.out, "you have not caught that pokemon")
		return nil
	}

	_, _ = fmt.Fprintf(r.out, "Name: %v\n", pokemon.Name)
	_, _ = fmt.Fprintf(r.out, "Height: %v\n", pokemon.Height)
	_, _ = fmt.Fprintf(r.out, "Weight: %v\n", pokemon.Weight)
	if len(pokemon.Stats) > 0 {
		_, _ = fmt.Fprintln(r.out, "Stats:")
		for _, stat := range pokemon.Stats {
			_, _ = fmt.Fprintf(r.out, "- %s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
	}
	if len(pokemon.Types) > 0 {
		_, _ = fmt.Fprintln(r.out, "Types:")
		for _, pokemonType := range pokemon.Types {
			_, _ = fmt.Fprintf(r.out, "- %v\n", pokemonType.Type.Name)
		}
	}

	return nil
}
