package registry

import (
	"go-bootcamp/interface/clients"
	"go-bootcamp/interface/controller"
	ci "go-bootcamp/usecase/clients"
)

func (r *registry) NewPokeAPIClient() ci.PokeAPIClient {
	return clients.NewPokeAPIClient(r.pokeAPIConfig)
}

func (r *registry) NewPokeAPIController() controller.PokeAPIController {
	return controller.NewPokeAPIController(r.NewPokeAPIClient(), r.NewPokemonInteractor())
}
