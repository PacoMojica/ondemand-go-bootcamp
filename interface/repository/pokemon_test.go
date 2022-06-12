package repository_test

import (
	"go-bootcamp/domain/model"
	"go-bootcamp/infrastructure/database"
	"go-bootcamp/interface/repository"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type mockDB struct {
	data [][]string
}

func newDB(v [][]string) database.DB {
	return &mockDB{v}
}

func (d *mockDB) Read() ([][]string, error) {
	return d.data, nil
}

func (d *mockDB) ConcurrentRead(f string, m, i int) ([][]string, error) {
	return d.data, nil
}

func (d *mockDB) Write(value []string) error {
	d.data = append(d.data, value)
	return nil
}
func (d *mockDB) WriteAll(values [][]string) error {
	d.data = append(d.data, values...)
	return nil
}

func TestPokemonRespositoryFindAll(t *testing.T) {
	var cases = map[string]struct {
		InitValues [][]string
		Expected   []model.Pokemon
	}{
		"2 pokemon": {
			InitValues: [][]string{
				{"132", "ditto", "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/132.png", "40", "3", "101", "ditto$-1$false$false$false$urban", "limber$https://pokeapi.co/api/v2/ability/7/$false,imposter$https://pokeapi.co/api/v2/ability/150/$true", "transform$https://pokeapi.co/api/v2/move/144/", "normal$https://pokeapi.co/api/v2/type/1/"},
			},
			Expected: []model.Pokemon{
				{
					ID:             132,
					Name:           "ditto",
					Sprites:        model.Sprites{FrontDefault: "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/132.png"},
					Weight:         40,
					Height:         3,
					BaseExperience: 101,
					Species: model.Species{
						Name:        "ditto",
						GenderRate:  -1,
						IsBaby:      false,
						IsLegendary: false,
						IsMythical:  false,
						Habitat:     "urban",
					},
					Abilities: []model.Ability{
						{
							Ability: model.AbilityInfo{
								Name: "limber",
								URL:  "https://pokeapi.co/api/v2/ability/7/",
							},
							IsHidden: false,
						},
						{
							Ability: model.AbilityInfo{
								Name: "imposter",
								URL:  "https://pokeapi.co/api/v2/ability/150/",
							},
							IsHidden: true,
						},
					},
					Moves: []model.Move{
						{Move: model.MoveInfo{
							Name: "transform",
							URL:  "https://pokeapi.co/api/v2/move/144/"}},
					},
					Types: []model.Type{
						{Type: model.TypeInfo{
							Name: "normal",
							URL:  "https://pokeapi.co/api/v2/type/1/",
						},
						},
					},
				},
			},
		},
		"0 pokemon": {
			InitValues: [][]string{},
			Expected:   *new([]model.Pokemon),
		},
	}

	for k, c := range cases {
		db := newDB(c.InitValues)
		r := repository.NewPokemonRepository(db)
		pokemon, err := r.FindAll()
		if err != nil {
			t.Fatalf(
				"[case %s]: pokemonRepository.FindAll(), unexpected error: '%s'",
				k, err,
			)
		}
		t.Log(pokemon)
		if !cmp.Equal(pokemon, c.Expected) {
			t.Errorf(
				"[case '%s']: pokemonRepository.FindAll()\nexpected\n-----\n%v\n-----\nbut got\n-----\n%v\n-----\ndiff\n-----\n%s\n",
				k, c.Expected, pokemon, cmp.Diff(pokemon, c.Expected))
		}
	}
}

func TestPokemonRespositoryFindById(t *testing.T) {
	var cases = map[string]struct {
		ID         uint
		InitValues [][]string
		Expected   model.Pokemon
	}{
		"2 pokemon": {
			ID: 132,
			InitValues: [][]string{
				{"132", "ditto", "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/132.png", "40", "3", "101", "ditto$-1$false$false$false$urban", "limber$https://pokeapi.co/api/v2/ability/7/$false,imposter$https://pokeapi.co/api/v2/ability/150/$true", "transform$https://pokeapi.co/api/v2/move/144/", "normal$https://pokeapi.co/api/v2/type/1/"},
			},
			Expected: model.Pokemon{
				ID:             132,
				Name:           "ditto",
				Sprites:        model.Sprites{FrontDefault: "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/132.png"},
				Weight:         40,
				Height:         3,
				BaseExperience: 101,
				Species: model.Species{
					Name:        "ditto",
					GenderRate:  -1,
					IsBaby:      false,
					IsLegendary: false,
					IsMythical:  false,
					Habitat:     "urban",
				},
				Abilities: []model.Ability{
					{
						Ability: model.AbilityInfo{
							Name: "limber",
							URL:  "https://pokeapi.co/api/v2/ability/7/",
						},
						IsHidden: false,
					},
					{
						Ability: model.AbilityInfo{
							Name: "imposter",
							URL:  "https://pokeapi.co/api/v2/ability/150/",
						},
						IsHidden: true,
					},
				},
				Moves: []model.Move{
					{Move: model.MoveInfo{
						Name: "transform",
						URL:  "https://pokeapi.co/api/v2/move/144/"}},
				},
				Types: []model.Type{
					{Type: model.TypeInfo{
						Name: "normal",
						URL:  "https://pokeapi.co/api/v2/type/1/",
					},
					},
				},
			},
		},
		"0 pokemon": {
			ID:         132,
			InitValues: [][]string{},
			Expected:   *new(model.Pokemon),
		},
	}

	for k, c := range cases {
		db := newDB(c.InitValues)
		r := repository.NewPokemonRepository(db)
		pokemon, err := r.FindById(c.ID)
		if err != nil {
			t.Fatalf(
				"[case %s]: pokemonRepository.FindById(%v), unexpected error: '%s'",
				k, c.ID, err,
			)
		}
		if !cmp.Equal(c.Expected, pokemon) {
			t.Errorf(
				"[case '%s']: pokemonRepository.FindById()\nexpected\n-----\n%v\n-----\nbut got\n-----\n%v\n-----\ndiff\n-----\n%s\n",
				k, c.Expected, pokemon, cmp.Diff(c.Expected, pokemon))
		}
	}
}

func TestPokemonRespositoryCreate(t *testing.T) {
	var cases = map[string]struct {
		InitValues [][]string
		Payload    model.Pokemon
		Expected   model.Pokemon
	}{
		"create lapras": {
			InitValues: [][]string{
				{"132", "ditto", "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/132.png", "40", "3", "101", "ditto$-1$false$false$false$urban", "limber$https://pokeapi.co/api/v2/ability/7/$false,imposter$https://pokeapi.co/api/v2/ability/150/$true", "transform$https://pokeapi.co/api/v2/move/144/", "normal$https://pokeapi.co/api/v2/type/1/"},
			},
			Payload: model.Pokemon{
				ID:   131,
				Name: "lapras",
				Sprites: model.Sprites{
					FrontDefault: "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/131.png",
				},
				Weight:         2200,
				Height:         25,
				BaseExperience: 187,
				Species: model.Species{
					Name:        "lapras",
					GenderRate:  0,
					IsBaby:      false,
					IsLegendary: false,
					IsMythical:  false,
					Habitat:     "",
				},
				Abilities: []model.Ability{
					{Ability: struct {
						Name string "json:\"name\""
						URL  string "json:\"url\""
					}{
						Name: "water-absorb",
						URL:  "https://pokeapi.co/api/v2/ability/11/",
					},
						IsHidden: false,
					},
				},
				Moves: []model.Move{
					{
						Move: struct {
							Name string "json:\"name\""
							URL  string "json:\"url\""
						}{
							Name: "hyper-beam",
							URL:  "https://pokeapi.co/api/v2/move/63/",
						},
					},
				},
				Types: []model.Type{
					{
						Type: struct {
							Name string "json:\"name\""
							URL  string "json:\"url\""
						}{
							Name: "water",
							URL:  "https://pokeapi.co/api/v2/type/11/",
						},
					},
					{
						Type: struct {
							Name string "json:\"name\""
							URL  string "json:\"url\""
						}{
							Name: "ice",
							URL:  "https://pokeapi.co/api/v2/type/15/",
						},
					},
				},
			},
			Expected: model.Pokemon{
				ID:   131,
				Name: "lapras",
				Sprites: model.Sprites{
					FrontDefault: "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/131.png",
				},
				Weight:         2200,
				Height:         25,
				BaseExperience: 187,
				Species: model.Species{
					Name:        "lapras",
					GenderRate:  0,
					IsBaby:      false,
					IsLegendary: false,
					IsMythical:  false,
					Habitat:     "",
				},
				Abilities: []model.Ability{
					{Ability: struct {
						Name string "json:\"name\""
						URL  string "json:\"url\""
					}{
						Name: "water-absorb",
						URL:  "https://pokeapi.co/api/v2/ability/11/",
					},
						IsHidden: false,
					},
				},
				Moves: []model.Move{
					{
						Move: struct {
							Name string "json:\"name\""
							URL  string "json:\"url\""
						}{
							Name: "hyper-beam",
							URL:  "https://pokeapi.co/api/v2/move/63/",
						},
					},
				},
				Types: []model.Type{
					{
						Type: struct {
							Name string "json:\"name\""
							URL  string "json:\"url\""
						}{
							Name: "water",
							URL:  "https://pokeapi.co/api/v2/type/11/",
						},
					},
					{
						Type: struct {
							Name string "json:\"name\""
							URL  string "json:\"url\""
						}{
							Name: "ice",
							URL:  "https://pokeapi.co/api/v2/type/15/",
						},
					},
				},
			},
		},
	}

	for k, c := range cases {
		db := newDB(c.InitValues)
		r := repository.NewPokemonRepository(db)
		err := r.Create(&c.Payload)
		if err != nil {
			t.Fatalf(
				"[case %s]: pokemonRepository.Create(%v), unexpected error: '%s'",
				k, c.Payload, err,
			)
		}
		if !cmp.Equal(c.Payload, c.Expected) {
			t.Errorf(
				"[case '%s']: pokemonRepository.Create(%v)\nexpected\n-----\n%v\n-----\nbut got\n-----\n%v\n-----\ndiff\n-----\n%s\n",
				k, c.Payload, c.Expected, c.Payload, cmp.Diff(c.Payload, c.Expected))
		}
	}
}
