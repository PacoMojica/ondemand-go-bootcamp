package repository

import (
	"fmt"
	"go-bootcamp/domain/model"
	"go-bootcamp/infrastructure/database"
	"go-bootcamp/usecase/repository"
)

type pokemonRepository struct {
	db database.DB
}

func NewPokemonRepository(db database.DB) repository.PokemonRepository {
	return &pokemonRepository{db}
}

func parsePokemon(r [][]string) ([]model.Pokemon, error) {
	var ps []model.Pokemon
	for _, record := range r {
		p := model.Pokemon{}

		if err := unmarshall(record, &p); err != nil {
			return nil, fmt.Errorf("Parsing pokemon data: %w", err)
		}

		ps = append(ps, p)
	}

	return ps, nil
}

func (pr *pokemonRepository) FindAll() ([]model.Pokemon, error) {
	r, err := pr.db.Read()
	if err != nil {
		return nil, err
	}
	ps, err := parsePokemon(r)
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (pr *pokemonRepository) FindById(ID uint) (p model.Pokemon, err error) {
	r, err := pr.db.Read()
	if err != nil {
		return
	}
	ps, err := parsePokemon(r)
	if err != nil {
		return
	}

	for _, item := range ps {
		if item.ID == ID {
			p = item
			break
		}
	}
	return
}

func (pr *pokemonRepository) Create(p *model.Pokemon) error {
	record := []string{}
	err := marshall(p, &record)
	if err != nil {
		return err
	}

	if err := pr.db.Write(record); err != nil {
		return err
	}

	return nil
}
