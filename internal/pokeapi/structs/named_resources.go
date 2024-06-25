package structs

type NamedApiResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type NamedApiResourceList struct {
	Count    int                `json:"count"`
	Next     *string            `json:"next"`
	Previous *string            `json:"previous"`
	Results  []NamedApiResource `json:"results"`
}

type Name struct {
	Name     string           `json:"name"`
	Language NamedApiResource `json:"language"`
}
