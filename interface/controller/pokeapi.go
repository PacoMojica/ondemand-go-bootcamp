package controller

import (
	"fmt"
	"go-bootcamp/config"
	"go-bootcamp/usecase/interactor"
	"math/rand"
	"net/http"
	"time"
)

type pokeAPIController struct {
	pokemonInteractor interactor.PokemonInteractor
}

type PokeAPIController interface {
	GetPokemon(res http.ResponseWriter, req *http.Request)
	GetPokemonFromIdentifier(res http.ResponseWriter, req *http.Request)
}

func NewPokeAPIController(pi interactor.PokemonInteractor) PokeAPIController {
	return &pokeAPIController{pi}
}

func (pa *pokeAPIController) GetPokemon(res http.ResponseWriter, req *http.Request) {
	if isValid := isMethodValid("GET", res, req); !isValid {
		return
	}

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	ID := (r.Intn(config.PokeAPI.TotalPokemon - 1)) + 1

	endpoint := fmt.Sprintf("%v/%v", config.PokeAPI.Endpoints.Pokemon, ID)
	apiRes, err := fetchPokeAPI(endpoint)
	if err != nil {
		writeError(fmt.Sprintf("could not fetch from the PokeAPI[%v]: %v", endpoint, err), res)
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

	endpoint := fmt.Sprintf("%v/%v", config.PokeAPI.Endpoints.Pokemon, identifier)
	apiRes, err := fetchPokeAPI(endpoint)
	if err != nil {
		writeError(fmt.Sprintf("could not fetch from the PokeAPI[%v]: %v", endpoint, err), res)
		return
	}

	data, err := pa.pokemonInteractor.Create(apiRes.Body)
	if err != nil {
		writeError(fmt.Sprintf("unable to save pokemon data: %v", err), res)
		return
	}

	writeJSON(data, res)
}
