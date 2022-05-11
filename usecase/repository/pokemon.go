package repository

import "go-bootcamp/domain/model"

type PokemonRepository interface {
	FindAll() ([]model.Pokemon, error)
	FindById(uint) (model.Pokemon, error)
	Create(*model.Pokemon) (*model.Pokemon, error)
}
