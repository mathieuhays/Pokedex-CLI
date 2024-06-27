package main

import (
	"errors"
	"fmt"
)

func commandPokemon(r *repl, args []string) error {
	if len(args) == 0 {
		return errors.New("please specify the name of a pokemon")
	}
	if len(args) > 1 {
		return errors.New("too many arguments")
	}

	pokemon, err := r.config.api.GetPokemon(args[0])
	if err != nil {
		return err
	}

	_, _ = fmt.Fprintf(r.out, "Pokemon: %v\n", pokemon.Name)
	_, _ = fmt.Fprintln(r.out, "=============")
	_, _ = fmt.Fprintf(r.out, "Base experience: %v\n", pokemon.BaseExperience)
	_, _ = fmt.Fprintln(r.out, "Stats: (base stat / effort)")
	if len(pokemon.Stats) == 0 {
		_, _ = fmt.Fprintln(r.out, "no stats available")
		return nil
	}
	for _, stat := range pokemon.Stats {
		_, _ = fmt.Fprintf(r.out, "- %v (%d/%d)\n", stat.Stat.Name, stat.BaseStat, stat.Effort)
	}

	return nil
}
