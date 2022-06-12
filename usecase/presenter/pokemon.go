package presenter

import (
	"io"

	"go-bootcamp/domain/model"
)

type PokemonPresenter interface {
	Marshall(any) ([]byte, error)
	Unmarshall(io.Reader) (model.Pokemon, error)
}
