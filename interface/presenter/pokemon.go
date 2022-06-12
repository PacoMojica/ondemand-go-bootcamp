package presenter

import (
	"encoding/json"
	"io"

	"go-bootcamp/domain/model"
	"go-bootcamp/usecase/presenter"
)

type pokemonPresenter struct{}

func NewPokemonPresenter() presenter.PokemonPresenter {
	return &pokemonPresenter{}
}

func (pp *pokemonPresenter) Marshall(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (pp *pokemonPresenter) Unmarshall(r io.Reader) (model.Pokemon, error) {
	var p model.Pokemon
	err := json.NewDecoder(r).Decode(&p)

	return p, err
}
