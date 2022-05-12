package registry

import (
	"go-bootcamp/interface/controller"
)

func (r *registry) NewPokeAPIController() controller.PokeAPIController {
	return controller.NewPokeAPIController(r.NewPokemonInteractor())
}
