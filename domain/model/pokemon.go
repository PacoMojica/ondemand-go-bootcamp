package model

type Ability struct {
	Name     string
	IsHidden bool
}

type Species struct {
	Name        string
	GenderRate  int
	IsBaby      bool
	IsLegendary bool
	IsMythical  bool
	Habitat     string
}

type Pokemon struct {
	ID             uint
	Name           string
	Image          string
	Weight         uint
	Height         uint
	BaseExperience uint
	Species        Species
	Abilities      []Ability
	Moves          []string
	Types          []string
}
