package pokeapi

type Pokemon struct {
	Id                     int
	Name                   string
	BaseExperience         int `json:"base_experience"`
	Height                 int
	IsDefault              bool `json:"is_default"`
	Order                  int
	Weight                 int
	Abilities              []PokemonAbility
	Forms                  []NamedApiResource
	GameIndices            []VersionGameIndex `json:"game_indices"`
	HeldItems              []PokemonHeldItem  `json:"held_items"`
	LocationAreaEncounters string             `json:"location_area_encounters"`
	Moves                  []PokemonMove
	PastTypes              []PokemonTypePast `json:"past_types"`
	Sprites                PokemonSprites
	Cries                  PokemonCries
	Species                NamedApiResource
	Stats                  []PokemonStat
	Types                  []PokemonType
}

type PokemonAbility struct {
	IsHidden bool `json:"is_hidden"`
	Slot     int
	Ability  NamedApiResource
}

type PokemonHeldItem struct {
	Item           NamedApiResource
	VersionDetails []PokemonHeldItemVersion `json:"version_details"`
}

type PokemonHeldItemVersion struct {
	Version NamedApiResource
	Rarity  int
}

type PokemonMove struct {
	Move                NamedApiResource
	VersionGroupDetails []PokemonMoveVersion `json:"version_group_details"`
}

type PokemonMoveVersion struct {
	MoveLearnMethod NamedApiResource `json:"move_learn_method"`
	VersionGroup    NamedApiResource `json:"version_group"`
	LevelLearnedAt  int              `json:"level_learned_at"`
}

type PokemonTypePast struct {
	Generation NamedApiResource
	Types      []PokemonType
}

// PokemonSprites more undocumented fields are available. @TODO maybe add support?
type PokemonSprites struct {
	FrontDefault     string `json:"front_default"`
	FrontShiny       string `json:"front_shiny"`
	FrontFemale      string `json:"front_female"`
	FrontShinyFemale string `json:"front_shiny_female"`
	BackDefault      string `json:"back_default"`
	BackShiny        string `json:"back_shiny"`
	BackFemale       string `json:"back_female"`
	BackShinyFemale  string `json:"back_shiny_female"`
}

type PokemonCries struct {
	Latest string
	Legacy string
}

type PokemonStat struct {
	Stat     NamedApiResource
	Effort   int
	BaseStat int `json:"base_stat"`
}

type PokemonType struct {
	Slot int
	Type NamedApiResource
}

type VersionGameIndex struct {
	GameIndex int `json:"game_index"`
	Version   NamedApiResource
}
