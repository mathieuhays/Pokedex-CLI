package pokeapi

import (
	"errors"
)

type NamedApiResource struct {
	Name string
	URL  string
}

type Name struct {
	Name     string
	Language NamedApiResource
}

type NamedApiResourceList struct {
	Count    int
	Next     *string
	Previous *string
	Results  []NamedApiResource
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
