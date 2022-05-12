package registry

import (
	"go-bootcamp/interface/controller"
	"go-bootcamp/interface/presenter"
	"go-bootcamp/interface/repository"
	"go-bootcamp/usecase/interactor"
	ip "go-bootcamp/usecase/presenter"
	ir "go-bootcamp/usecase/repository"
)

func (r *registry) NewPokemonPresenter() ip.PokemonPresenter {
	return presenter.NewPokemonPresenter()
}

func (r *registry) NewPokemonRepository() ir.PokemonRepository {
	return repository.NewPokemonRepository(r.db)
}

func (r *registry) NewPokemonInteractor() interactor.PokemonInteractor {
	return interactor.NewPokemonInteractor(r.NewPokemonRepository(), r.NewPokemonPresenter())
}

func (r *registry) NewPokemonController() controller.PokemonController {
	return controller.NewPokemonController(r.NewPokemonInteractor())
}
