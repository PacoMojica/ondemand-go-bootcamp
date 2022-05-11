package registry

import (
	"go-bootcamp/interface/controller"
	"go-bootcamp/interface/repository"
	"go-bootcamp/usecase/interactor"
	ir "go-bootcamp/usecase/repository"
)

func (r *registry) NewPokemonRepository() ir.PokemonRepository {
	return repository.NewPokemonRepository(r.db)
}

func (r *registry) NewPokemonInteractor() interactor.PokemonInteractor {
	return interactor.NewPokemonInteractor(r.NewPokemonRepository())
}

func (r *registry) NewPokemonController() controller.PokemonController {
	return controller.NewPokemonController(r.NewPokemonInteractor())
}
