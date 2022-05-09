package interactor

import (
	"go-bootcamp/domain/model"
	"go-bootcamp/usecase/repository"
)

type pokemonInteractor struct {
	PokemonRepository repository.PokemonRepository
}

type PokemonInteractor interface {
	GetAll() ([]model.Pokemon, error)
	GetById(uint) (model.Pokemon, error)
	Create(p *model.Pokemon) error
}

func New(r repository.PokemonRepository) PokemonInteractor {
	return &pokemonInteractor{r}
}

func (pi *pokemonInteractor) GetAll() ([]model.Pokemon, error) {
	p, err := pi.PokemonRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (pi *pokemonInteractor) GetById(ID uint) (p model.Pokemon, err error) {
	return pi.PokemonRepository.FindById(ID)
}

func (pi *pokemonInteractor) Create(p *model.Pokemon) error {
	p, err := pi.PokemonRepository.Create(p)
	if err != nil {
		return err
	}

	return nil
}
