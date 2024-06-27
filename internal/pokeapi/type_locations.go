package pokeapi

type LocationArea struct {
	Id                   int
	Name                 string
	GameIndex            int                   `json:"game_index"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Location             NamedApiResource
	Names                []Name
	PokemonEncounters    []PokemonEncounter `json:"pokemon_encounters"`
}
