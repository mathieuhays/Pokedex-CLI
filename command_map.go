package main

import (
	"fmt"
)

func renderLocation(r *repl, locationURL string) error {
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
	url := r.config.api.GetLocationResourceURL()

	if r.config.nextLocationURL != nil {
		url = *r.config.nextLocationURL
	}

	return renderLocation(r, url)
}

func commandPreviousMap(r *repl, args []string) error {
	if r.config.previousLocationURL == nil {
		_, _ = fmt.Fprintln(r.out, "No previous location to show")
		return nil
	}

	return renderLocation(r, *r.config.previousLocationURL)
}
