package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"go-bootcamp/usecase/interactor"
)

type pokemonController struct {
	interactor interactor.PokemonInteractor
}

type PokemonController interface {
	// Returns all pokemon stored in the DB
	GetPokemon(res http.ResponseWriter, req *http.Request)
	// Returns the pokemon that matches the ID or null
	GetPokemonById(res http.ResponseWriter, req *http.Request)
	// Creates a pokemon using the post data
	CreatePokemon(res http.ResponseWriter, req *http.Request)
	// Reads the database concurrently
	ConcurrentPokemon(res http.ResponseWriter, req *http.Request)
}

// Returns an instance of the pokemon controller
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

func (pc *pokemonController) ConcurrentPokemon(res http.ResponseWriter, req *http.Request) {
	if isValid := isMethodValid("GET", res, req); !isValid {
		return
	}
	// type: Only support "odd" or "even"
	// items: Is an Int and is the amount of valid items you need to display as a response
	// items_per_workers: I
	params := req.URL.Query()
	filter := params.Get("type")
	if filter != "even" && filter != "odd" {
		writeError(fmt.Sprintf("invalid value [type]: '%v'", filter), res)
		return
	}
	items := params.Get("items")
	maxItems, err := strconv.ParseInt(items, 10, 64)
	if err != nil {
		writeError(fmt.Sprintf("invalid value [items:%v]: %v", items, err), res)
		return
	}

	perWorker := params.Get("items_per_workers")
	itemsPerWorker, err := strconv.ParseInt(perWorker, 10, 64)
	if err != nil {
		writeError(fmt.Sprintf("invalid value [items_per_workers:%v]: %v", perWorker, err), res)
		return
	}

	p, err := pc.interactor.GetAllConcurrently(filter, int(maxItems), int(itemsPerWorker))
	if err != nil {
		writeError(fmt.Sprintf("could not fetch all pokemon: %v", err), res)
		return
	}

	writeJSON(p, res)
}
