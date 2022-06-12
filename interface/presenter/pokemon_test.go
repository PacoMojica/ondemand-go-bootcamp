package presenter_test

import (
	"bytes"
	"flag"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"

	"go-bootcamp/domain/model"
	"go-bootcamp/interface/presenter"
)

var update = flag.Bool("update", false, "update golden files")

var Pokemon = model.Pokemon{
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
}

func readFile(name, k string, t *testing.T) []byte {
	bytes, err := ioutil.ReadFile(name)
	if err != nil {
		t.Fatalf(
			"[case '%s']: unexpected error reading file '%s': '%s'",
			k, name, err)
	}
	return bytes
}

func writeFile(name, k string, content []byte, t *testing.T) {
	err := ioutil.WriteFile(name, content, 0644)
	if err != nil {
		t.Fatalf(
			"[case '%s']: unexpected error writing file (%s): '%s'",
			k, name, err)
	}
}

func getGoldenFile(name, k string, content []byte, t *testing.T) []byte {
	golden := filepath.Join("testdata", name+".golden")
	if *update {
		writeFile(golden, k, content, t)
	}
	expected := readFile(golden, k, t)
	return expected
}

func TestPokemonPresenterMarshall(t *testing.T) {
	var cases = map[string]struct {
		Value  model.Pokemon
		Golden string
	}{
		"marshall pokemon": {Value: Pokemon, Golden: "lapras"},
	}

	for k, c := range cases {
		data, err := presenter.NewPokemonPresenter().Marshall(Pokemon)
		if err != nil {
			t.Fatalf(
				"[case %s]: pokemonPresenter.Marshall(%s): '%s'",
				k, Pokemon.Name, err,
			)
		}
		expected := getGoldenFile(c.Golden, k, data, t)

		if !bytes.Equal(data, expected) {
			t.Errorf(
				"[case '%s']: pokemonPresenter.Marshall('%s')\nexpected\n-----\n%s\n-----\nbut got\n-----\n%s\n-----\n",
				k, c.Value.Name, expected, data)
		}
	}
}

func TestPokemonPresenterUnmarshall(t *testing.T) {
	var cases = map[string]struct {
		Expected model.Pokemon
		JSON     string
	}{
		"unmarshall pokemon": {Expected: Pokemon, JSON: "lapras"},
	}

	for k, c := range cases {
		data := readFile("testdata/"+c.JSON+".json", k, t)
		reader := bytes.NewReader(data)
		pokemon, err := presenter.NewPokemonPresenter().Unmarshall(reader)
		if err != nil {
			t.Fatalf(
				"[case %s]: pokemonPresenter.Marshall(%s): '%s'",
				k, Pokemon.Name, err,
			)
		}

		if !cmp.Equal(c.Expected, pokemon) {
			t.Errorf(
				"[case '%s']: pokemonPresenter.Unmarshall('%s')\nexpected\n-----\n%v\n-----\nbut got\n-----\n%v\n-----\ndiff\n-----\n%s",
				k, pokemon.Name, c.Expected, pokemon, cmp.Diff(c.Expected, pokemon))
		}
	}
}
