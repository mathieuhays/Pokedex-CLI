package main

import "fmt"

func commandList(r *repl, args []string) error {
	pokemons := r.config.pokedex.GetAll()

	if len(pokemons) == 0 {
		_, _ = fmt.Fprintln(r.out, "you have no pokemons at the moment")
		return nil
	}

	for _, pokemon := range pokemons {
		_, _ = fmt.Fprintf(r.out, "- %v (after %d attemtps)\n", pokemon.Name, r.config.pokedex.GetNumberOfAttempts(pokemon.Name))
	}

	return nil
}
