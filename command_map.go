package main

import (
	"fmt"
	"github.com/mathieuhays/pokedex-cli/internal/pokeapi"
)

func renderLocation(r *repl, resource *pokeapi.NamedApiResourceList) error {
	for _, result := range resource.Results {
		_, _ = fmt.Fprintf(r.out, "%s\n", result.Name)
	}

	return nil
}

func commandMap(r *repl, args []string) error {
	var resource pokeapi.NamedApiResourceList
	var err error

	if r.config.lastLocation != nil {
		resource, err = r.config.lastLocation.NextPage()
	} else {
		resource, err = r.config.api.ListLocations()
	}

	if err != nil {
		return err
	}

	r.config.lastLocation = &resource
	return renderLocation(r, &resource)
}

func commandPreviousMap(r *repl, args []string) error {
	var resource pokeapi.NamedApiResourceList
	var err error

	if r.config.lastLocation != nil {
		resource, err = r.config.lastLocation.PreviousPage()
	} else {
		resource, err = r.config.api.ListLocations()
	}

	if err != nil {
		return err
	}

	r.config.lastLocation = &resource
	return renderLocation(r, &resource)
}
