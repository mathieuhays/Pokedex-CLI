package main

import (
	"fmt"
)

func renderLocation(r *repl, locationURL *string) error {
	resource, err := r.config.api.ListLocations(locationURL)
	if err != nil {
		return err
	}

	r.config.previousLocationURL = resource.Previous
	r.config.nextLocationURL = resource.Next

	for _, result := range resource.Results {
		_, _ = fmt.Fprintf(r.out, "%s\n", result.Name)
	}

	return nil
}

func commandMap(r *repl, args []string) error {
	return renderLocation(r, r.config.nextLocationURL)
}

func commandPreviousMap(r *repl, args []string) error {
	return renderLocation(r, r.config.previousLocationURL)
}
