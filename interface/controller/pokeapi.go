package controller

import (
	"fmt"
	"go-bootcamp/usecase/clients"
	"go-bootcamp/usecase/interactor"
	"net/http"
)

type pokeAPIController struct {
	client            clients.PokeAPIClient
	pokemonInteractor interactor.PokemonInteractor
}

type PokeAPIController interface {
	GetPokemon(res http.ResponseWriter, req *http.Request)
	GetPokemonFromIdentifier(res http.ResponseWriter, req *http.Request)
}

func NewPokeAPIController(c clients.PokeAPIClient, pi interactor.PokemonInteractor) PokeAPIController {
	return &pokeAPIController{c, pi}
}

func (pa *pokeAPIController) GetPokemon(res http.ResponseWriter, req *http.Request) {
	if isValid := isMethodValid("GET", res, req); !isValid {
		return
	}

	apiRes, err := pa.client.FetchRandomPokemon()
	if err != nil {
		writeError(fmt.Sprintf("could not fetch random pokemon from the PokeAPI: %v", err), res)
		return
	}

	data, err := pa.pokemonInteractor.Create(apiRes.Body)
	if err != nil {
		writeError(fmt.Sprintf("unable to save pokemon data: %v", err), res)
		return
	}

	writeJSON(data, res)
}

func (pa *pokeAPIController) GetPokemonFromIdentifier(res http.ResponseWriter, req *http.Request) {
	if isValid := isMethodValid("GET", res, req); !isValid {
		return
	}

	paths, ok := getPaths(res, req)
	if !ok {
		return
	}

	identifier, ok := getPokemonIdentifier(paths, res)
	if !ok {
		return
	}

	apiRes, err := pa.client.FetchPokemon(identifier)
	if err != nil {
		writeError(fmt.Sprintf("could not fetch from the PokeAPI[%v]: %v", identifier, err), res)
		return
	}

	data, err := pa.pokemonInteractor.Create(apiRes.Body)
	if err != nil {
		writeError(fmt.Sprintf("unable to save pokemon data: %v", err), res)
		return
	}

	writeJSON(data, res)
}
