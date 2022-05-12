package model

type Sprites struct {
	FrontDefault string `json:"front_default"`
}

type Species struct {
	Name        string `json:"name"`
	GenderRate  int    `json:"gender_rate"`
	IsBaby      bool   `json:"is_baby"`
	IsLegendary bool   `json:"is_legendary"`
	IsMythical  bool   `json:"is_mythical"`
	Habitat     string `json:"habitad"`
}

type Ability struct {
	Ability struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"ability"`
	IsHidden bool `json:"is_hidden"`
}

type Move struct {
	Move struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"move"`
}

type Type struct {
	Type struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"type"`
}

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
