package controller

import (
	"fmt"
	"net/http"

	"go-bootcamp/usecase/interactor"
)

type pokemonController struct {
	interactor interactor.PokemonInteractor
}

type PokemonController interface {
	GetPokemon(res http.ResponseWriter, req *http.Request)
	GetPokemonById(res http.ResponseWriter, req *http.Request)
	CreatePokemon(res http.ResponseWriter, req *http.Request)
}

func NewPokemonController(pi interactor.PokemonInteractor) PokemonController {
	return &pokemonController{pi}
}

func (pc *pokemonController) GetPokemon(res http.ResponseWriter, req *http.Request) {
	if isValid := isMethodValid("GET", res, req); !isValid {
		return
	}

	p, err := pc.interactor.GetAll()
	if err != nil {
		writeError(fmt.Sprintf("could not fetch all pokemon: %v", err), res)
		return
	}

	writeJSON(p, res)
}

func (pc *pokemonController) GetPokemonById(res http.ResponseWriter, req *http.Request) {
	if ok := isMethodValid("GET", res, req); !ok {
		return
	}

	paths, ok := getPaths(res, req)
	if !ok {
		return
	}

	ID, ok := getPathID(paths, res)
	if !ok {
		return
	}

	p, err := pc.interactor.GetById(ID)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		writeError(fmt.Sprintf("could not fetch the pokemon with ID [%v] %v", ID, err), res)
		return
	}

	writeJSON(p, res)
}

func (pc *pokemonController) CreatePokemon(res http.ResponseWriter, req *http.Request) {
	if isValid := isMethodValid("POST", res, req); !isValid {
		return
	}

	data, err := pc.interactor.Create(req.Body)
	if err != nil {
		writeError(fmt.Sprintf("unable to save pokemon data: %v", err), res)
		return
	}

	writeJSON(data, res)
}
