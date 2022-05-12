package model

type Ability struct {
	Name     string `json:"name"`
	IsHidden bool   `json:"is_hidden"`
}

type Species struct {
	Name        string `json:"name"`
	GenderRate  int    `json:"gender_rate"`
	IsBaby      bool   `json:"is_baby"`
	IsLegendary bool   `json:"is_legendary"`
	IsMythical  bool   `json:"is_mythical"`
	Habitat     string `json:"habitad"`
}

type Pokemon struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	Image          string    `json:"image"`
	Weight         uint      `json:"weight"`
	Height         uint      `json:"height"`
	BaseExperience uint      `json:"base_experience"`
	Species        Species   `json:"species"`
	Abilities      []Ability `json:"abilities"`
	Moves          []string  `json:"moves"`
	Types          []string  `json:"types"`
}
