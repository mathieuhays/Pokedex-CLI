package pokeapi

import (
	"errors"
)

type NamedApiResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Name struct {
	Name     string           `json:"name"`
	Language NamedApiResource `json:"language"`
}

type NamedApiResourceList struct {
	Count    int                `json:"count"`
	Next     *string            `json:"next"`
	Previous *string            `json:"previous"`
	Results  []NamedApiResource `json:"results"`
	Client   *Client
}

func (l *NamedApiResourceList) NextPage() (NamedApiResourceList, error) {
	if l.Next == nil {
		return NamedApiResourceList{}, errors.New("no new page to load")
	}

	return l.Client.listLocationsWithURL(*l.Next)
}

func (l *NamedApiResourceList) PreviousPage() (NamedApiResourceList, error) {
	if l.Previous == nil {
		return NamedApiResourceList{}, errors.New("no previous page to load")
	}

	return l.Client.listLocationsWithURL(*l.Previous)
}
