package repository

import (
	"fmt"
	"go-bootcamp/domain/model"
	"go-bootcamp/infrastructure/database"
	"go-bootcamp/usecase/repository"
	"strconv"
	"strings"
)

type pokemonRepository struct {
	db database.DB
}

func New(db database.DB) repository.PokemonRepository {
	return &pokemonRepository{db}
}

func parsePokemon(r [][]string) ([]model.Pokemon, error) {
	var ps []model.Pokemon
	for _, record := range r {
		p := model.Pokemon{}

		if err := unmarshallRecord(record, &p); err != nil {
			return nil, fmt.Errorf("Parsing pokemon data: %w", err)
		}

		ps = append(ps, p)
	}

	return ps, nil
}

func parseRecord(p *model.Pokemon) []string {
	// TODO: create inverse of unmarshallRecord
	species := strings.Join([]string{
		p.Species.Name,
		strconv.FormatInt(int64(p.Species.GenderRate), 10),
		strconv.FormatBool(bool(p.Species.IsBaby)),
		strconv.FormatBool(bool(p.Species.IsLegendary)),
		strconv.FormatBool(bool(p.Species.IsMythical)),
		p.Species.Habitat,
	}, "$")
	aSlice := []string{}
	for _, a := range p.Abilities {
		aSlice = append(aSlice, fmt.Sprintf("%v$%v", a.Name, a.IsHidden))
	}
	abilities := fmt.Sprintf("%v", strings.Join(aSlice, ","))
	moves := fmt.Sprintf("%v", strings.Join(p.Moves, ","))
	types := fmt.Sprintf("%v", strings.Join(p.Types, ","))

	return []string{
		strconv.FormatUint(uint64(p.ID), 10),
		p.Name,
		p.Image,
		strconv.FormatUint(uint64(p.Weight), 10),
		strconv.FormatUint(uint64(p.Height), 10),
		strconv.FormatUint(uint64(p.BaseExperience), 10),
		species,
		abilities,
		moves,
		types,
	}
}

func (pr *pokemonRepository) FindAll() ([]model.Pokemon, error) {
	r, err := pr.db.Read()
	if err != nil {
		return nil, err
	}
	p, err := parsePokemon(r)
	if err != nil {
		return nil, err
	}

	return p, nil
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

func (pr *pokemonRepository) Create(p *model.Pokemon) (*model.Pokemon, error) {
	r := parseRecord(p)
	if err := pr.db.Write(r); err != nil {
		return nil, err
	}

	return p, nil
}
