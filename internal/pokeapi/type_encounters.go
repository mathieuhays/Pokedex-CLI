package pokeapi

type EncounterMethodRate struct {
	EncounterMethod NamedApiResource          `json:"encounter_method"`
	VersionDetails  []EncounterVersionDetails `json:"version_details"`
}

type EncounterVersionDetails struct {
	Rate    int
	Version NamedApiResource
}

type VersionEncounterDetail struct {
	Version          NamedApiResource
	MaxChance        int         `json:"max_chance"`
	EncounterDetails []Encounter `json:"encounter_details"`
}

type Encounter struct {
	MinLevel        int                `json:"min_level"`
	MaxLevel        int                `json:"max_level"`
	ConditionValues []NamedApiResource `json:"condition_values"`
	Chance          int
	Method          NamedApiResource
}

type PokemonEncounter struct {
	Pokemon        NamedApiResource
	VersionDetails []VersionEncounterDetail `json:"version_details"`
}
