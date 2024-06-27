package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(r *repl, args []string) error {
	if len(args) == 0 {
		return errors.New("please specify a pokemon")
	}
	if len(args) > 1 {
		return errors.New("too many arguments")
	}

	pokemon, err := r.config.api.GetPokemon(args[0])
	if err != nil {
		return err
	}

	if r.config.pokedex.Has(pokemon.Name) {
		_, _ = fmt.Fprintln(r.out, "you already have that pokemon")
		return nil
	}

	r.config.pokedex.RegisterAttempt(pokemon.Name)
	_, _ = fmt.Fprintf(r.out, "Throwing a pokeball at %v", pokemon.Name)

	for i := 0; i < 4; i++ {
		time.Sleep(time.Millisecond * 400)
		_, _ = fmt.Fprint(r.out, ".")
	}

	_, _ = fmt.Fprint(r.out, "\n")

	if isCaught(pokemon.BaseExperience) {
		err = r.config.pokedex.Add(pokemon)
		if err != nil {
			return err
		}

		_, _ = fmt.Fprintf(r.out, "Well done! You caught %v! This pokemon has been added to your pokedex.\n", pokemon.Name)
		return nil
	}

	_, _ = fmt.Fprintf(r.out, "%v escaped!\n", pokemon.Name)

	return nil
}

func isCaught(experience int) bool {
	return rand.Intn(experience) < 40
}

/*
pikachu: 112
bulbasaur: 64
rattata: 51
nidoran-f 55
jigglypuff: 95
vulpix: 60
mew: 300
electrode: 172
ditto: 101
vaporeon: 184
snorlax: 189
mewtwo: 340
*/
