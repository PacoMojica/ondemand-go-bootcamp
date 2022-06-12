package presenter

import (
	"go-bootcamp/domain/model"
	"io"
)

type PokemonPresenter interface {
	Marshall(any) ([]byte, error)
	Unmarshall(io.Reader) (model.Pokemon, error)
}
