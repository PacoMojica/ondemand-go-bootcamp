package interactor_test

import (
	"bytes"
	"flag"
	"go-bootcamp/infrastructure/database"
	"go-bootcamp/interface/presenter"
	"go-bootcamp/interface/repository"
	"go-bootcamp/usecase/interactor"
	"io/ioutil"
	"path/filepath"
	"testing"
)

var update = flag.Bool("update", false, "update golden files")

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

func TestGetAll(t *testing.T) {
	var cases = map[string]struct {
		Golden     string
		InitValues [][]string
	}{
		"get only ditto": {
			Golden: "getall",
			InitValues: [][]string{
				{"132", "ditto", "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/132.png", "40", "3", "101", "ditto$-1$false$false$false$urban", "limber$https://pokeapi.co/api/v2/ability/7/$false,imposter$https://pokeapi.co/api/v2/ability/150/$true", "transform$https://pokeapi.co/api/v2/move/144/", "normal$https://pokeapi.co/api/v2/type/1/"},
			},
		},
	}

	for k, c := range cases {
		db := newDB(c.InitValues)
		r := repository.NewPokemonRepository(db)
		p := presenter.NewPokemonPresenter()
		i := interactor.NewPokemonInteractor(r, p)
		d, err := i.GetAll()
		if err != nil {
			t.Fatalf(
				"[case %s]: pokemonInteractor.GetAll(): '%s'",
				k, err,
			)
		}
		expected := getGoldenFile(c.Golden, k, d, t)
		if !bytes.Equal(d, expected) {
			t.Errorf(
				"[case '%s']: pokemonInteractor.GetAll()\nexpected\n-----\n%s\n-----\nbut got\n-----\n%s\n-----\n",
				k, expected, d)
		}
	}
}

func TestGetById(t *testing.T) {
	var cases = map[string]struct {
		Golden     string
		InitValues [][]string
		ID         uint
	}{
		"get only ditto": {
			Golden: "getbyid",
			InitValues: [][]string{
				{"132", "ditto", "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/132.png", "40", "3", "101", "ditto$-1$false$false$false$urban", "limber$https://pokeapi.co/api/v2/ability/7/$false,imposter$https://pokeapi.co/api/v2/ability/150/$true", "transform$https://pokeapi.co/api/v2/move/144/", "normal$https://pokeapi.co/api/v2/type/1/"},
			},
			ID: 132,
		},
		"return null if ID is not found": {
			Golden: "getbyid-null",
			InitValues: [][]string{
				{"132", "ditto", "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/132.png", "40", "3", "101", "ditto$-1$false$false$false$urban", "limber$https://pokeapi.co/api/v2/ability/7/$false,imposter$https://pokeapi.co/api/v2/ability/150/$true", "transform$https://pokeapi.co/api/v2/move/144/", "normal$https://pokeapi.co/api/v2/type/1/"},
			},
			ID: 131,
		},
	}

	for k, c := range cases {
		db := newDB(c.InitValues)
		r := repository.NewPokemonRepository(db)
		p := presenter.NewPokemonPresenter()
		i := interactor.NewPokemonInteractor(r, p)
		d, err := i.GetById(c.ID)
		if err != nil {
			t.Fatalf(
				"[case %s]: pokemonInteractor.GetById(%d): '%s'",
				k, c.ID, err,
			)
		}
		expected := getGoldenFile(c.Golden, k, d, t)
		if !bytes.Equal(d, expected) {
			t.Errorf(
				"[case '%s']: pokemonInteractor.GetById(%d)\nexpected\n-----\n%s\n-----\nbut got\n-----\n%s\n-----\n",
				k, c.ID, expected, d)
		}
	}
}

func TestCreate(t *testing.T) {
	var cases = map[string]struct {
		Golden     string
		InitValues [][]string
		JSON       string
	}{
		"create lapras": {
			Golden: "create",
			InitValues: [][]string{
				{"132", "ditto", "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/132.png", "40", "3", "101", "ditto$-1$false$false$false$urban", "limber$https://pokeapi.co/api/v2/ability/7/$false,imposter$https://pokeapi.co/api/v2/ability/150/$true", "transform$https://pokeapi.co/api/v2/move/144/", "normal$https://pokeapi.co/api/v2/type/1/"},
			},
			JSON: "lapras",
		},
	}

	for k, c := range cases {
		db := newDB(c.InitValues)
		r := repository.NewPokemonRepository(db)
		p := presenter.NewPokemonPresenter()
		i := interactor.NewPokemonInteractor(r, p)
		payload := readFile(filepath.Join("testdata", c.JSON+".json"), k, t)
		reader := bytes.NewBuffer(payload)
		d, err := i.Create(reader)
		if err != nil {
			t.Fatalf(
				"[case %s]: pokemonInteractor.Create(%s): '%s'",
				k, payload, err,
			)
		}
		expected := getGoldenFile(c.Golden, k, d, t)
		if !bytes.Equal(d, expected) {
			t.Errorf(
				"[case '%s']: pokemonInteractor.Create(%s)\nexpected\n-----\n%s\n-----\nbut got\n-----\n%s\n-----\n",
				k, payload, expected, d)
		}
	}
}
