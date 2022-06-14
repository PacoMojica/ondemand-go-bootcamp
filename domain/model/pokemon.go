package model

// The sprites of a pokemon
type Sprites struct {
	FrontDefault string `json:"front_default"`
}

// A pokemon species
type Species struct {
	Name        string `json:"name"`
	GenderRate  int    `json:"gender_rate"`
	IsBaby      bool   `json:"is_baby"`
	IsLegendary bool   `json:"is_legendary"`
	IsMythical  bool   `json:"is_mythical"`
	Habitat     string `json:"habitad"`
}

// Attributes of the pokemon ability
type AbilityInfo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// A pokemon ability
type Ability struct {
	Ability  AbilityInfo `json:"ability"`
	IsHidden bool        `json:"is_hidden"`
}

// Attributes of the pokemon move
type MoveInfo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// A pokemon move
type Move struct {
	Move MoveInfo `json:"move"`
}

// Attributes of the pokemon type
type TypeInfo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// A pokemon type
type Type struct {
	Type TypeInfo `json:"type"`
}

// Represents a pokemon
type Pokemon struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	Sprites        Sprites   `json:"sprites"`
	Weight         uint      `json:"weight"`
	Height         uint      `json:"height"`
	BaseExperience uint      `json:"base_experience"`
	Species        Species   `json:"species"`
	Abilities      []Ability `json:"abilities"`
	Moves          []Move    `json:"moves"`
	Types          []Type    `json:"types"`
}
