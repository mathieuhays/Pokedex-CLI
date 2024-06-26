package pokeapi

type EncounterMethodRate struct {
	EncounterMethod NamedApiResource          `json:"encounter_method"`
	VersionDetails  []EncounterVersionDetails `json:"version_details"`
}

type EncounterVersionDetails struct {
	Rate    int              `json:"rate"`
	Version NamedApiResource `json:"version"`
}

type VersionEncounterDetail struct {
	Version          NamedApiResource `json:"version"`
	MaxChance        int              `json:"max_chance"`
	EncounterDetails []Encounter      `json:"encounter_details"`
}

type Encounter struct {
	MinLevel        int                `json:"min_level"`
	MaxLevel        int                `json:"max_level"`
	ConditionValues []NamedApiResource `json:"condition_values"`
	Chance          int                `json:"chance"`
	Method          NamedApiResource   `json:"method"`
}

type PokemonEncounter struct {
	Pokemon        NamedApiResource         `json:"pokemon"`
	VersionDetails []VersionEncounterDetail `json:"version_details"`
}
