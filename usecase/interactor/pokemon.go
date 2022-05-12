package interactor

import (
	"go-bootcamp/usecase/presenter"
	"go-bootcamp/usecase/repository"
	"io"
)

type pokemonInteractor struct {
	PokemonRepository repository.PokemonRepository
	PokemonPresenter  presenter.PokemonPresenter
}

type PokemonInteractor interface {
	GetAll() ([]byte, error)
	GetById(uint) ([]byte, error)
	Create(r io.Reader) ([]byte, error)
}

func NewPokemonInteractor(r repository.PokemonRepository, p presenter.PokemonPresenter) PokemonInteractor {
	return &pokemonInteractor{r, p}
}

func (pi *pokemonInteractor) GetAll() ([]byte, error) {
	p, err := pi.PokemonRepository.FindAll()
	if err != nil {
		return nil, err
	}

	data, err := pi.PokemonPresenter.Marshall(p)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (pi *pokemonInteractor) GetById(ID uint) ([]byte, error) {
	p, err := pi.PokemonRepository.FindById(ID)
	if err != nil {
		return nil, err
	}

	var data []byte
	if p.ID != uint(ID) {
		data, err = pi.PokemonPresenter.Marshall(nil)
	} else {
		data, err = pi.PokemonPresenter.Marshall(p)
	}

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (pi *pokemonInteractor) Create(r io.Reader) ([]byte, error) {
	p, err := pi.PokemonPresenter.Unmarshall(r)
	if err != nil {
		return nil, err
	}

	err = pi.PokemonRepository.Create(&p)
	if err != nil {
		return nil, err
	}

	data, err := pi.PokemonPresenter.Marshall(p)
	if err != nil {
		return nil, err
	}

	return data, nil
}
