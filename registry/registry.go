package registry

import (
	"go-bootcamp/infrastructure/database"
	"go-bootcamp/interface/controller"
)

type registry struct {
	db database.DB
}

type Registry interface {
	NewAppController() controller.AppController
}

func New(db database.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		Pokemon: r.NewPokemonController(),
	}
}
