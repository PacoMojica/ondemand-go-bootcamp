package registry

import (
	"go-bootcamp/infrastructure/database"
	"go-bootcamp/interface/controller"
	"go-bootcamp/usecase/clients"
)

type registry struct {
	db            database.DB
	pokeAPIConfig clients.PokeAPIConfig
}

type Registry interface {
	NewAppController() controller.AppController
}

func New(db database.DB, pac clients.PokeAPIConfig) Registry {
	return &registry{db, pac}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		Pokemon: r.NewPokemonController(),
		PokeAPI: r.NewPokeAPIController(),
	}
}
