package interactor

import (
	"io"

	"go-bootcamp/usecase/presenter"
	"go-bootcamp/usecase/repository"
)

type pokemonInteractor struct {
	PokemonRepository repository.PokemonRepository
	PokemonPresenter  presenter.PokemonPresenter
}

type PokemonInteractor interface {
	// Returns all the pokemon in the DB
	GetAll() ([]byte, error)
	// Returns the pokemon that matches the ID
	GetById(uint) ([]byte, error)
	// Creates a new pokemon
	Create(r io.Reader) ([]byte, error)
	// Reads the DB concurrently and returns all the pokemon
	GetAllConcurrently(filter string, maxItems int, itemsPerWorker int) ([]byte, error)
}

// Returns a new instance of the pokemon interactor
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

	found, err := pi.PokemonRepository.FindById(p.ID)
	if err != nil {
		return nil, err
	}

	if found.ID != p.ID {
		err = pi.PokemonRepository.Create(&p)
		if err != nil {
			return nil, err
		}
	}

	data, err := pi.PokemonPresenter.Marshall(p)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (pi *pokemonInteractor) GetAllConcurrently(
	filter string, maxItems int, itemsPerWorker int,
) ([]byte, error) {
	p, err := pi.PokemonRepository.FindAllConcurrently(filter, maxItems, itemsPerWorker)
	if err != nil {
		return nil, err
	}

	data, err := pi.PokemonPresenter.Marshall(p)
	if err != nil {
		return nil, err
	}

	return data, nil
}
