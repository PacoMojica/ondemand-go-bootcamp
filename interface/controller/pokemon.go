package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"go-bootcamp/domain/model"
	"go-bootcamp/usecase/interactor"
)

type pokemonController struct {
	interactor interactor.PokemonInteractor
}

type PokemonController interface {
	GetPokemon(w http.ResponseWriter, req *http.Request)
	GetPokemonById(w http.ResponseWriter, req *http.Request)
	CreatePokemon(w http.ResponseWriter, req *http.Request)
}

func New(pi interactor.PokemonInteractor) PokemonController {
	return &pokemonController{pi}
}

func (pc *pokemonController) GetPokemon(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		message := fmt.Sprintf("unsupported method [%v]", req.Method)
		log.Println(message)
		fmt.Fprint(w, message)
		return
	}

	p, err := pc.interactor.GetAll()
	if err != nil {
		message := fmt.Sprintf("could not fetch all pokemon: %v", err)
		log.Println(message)
		fmt.Fprint(w, message)
		return
	}
	data, err := json.Marshal(p)
	if err != nil {
		message := fmt.Sprintf("unable to marshall response: %v", err)
		log.Println(message)
		fmt.Fprint(w, message)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (pc *pokemonController) GetPokemonById(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		message := fmt.Sprintf("unsupported method [%v]", req.Method)
		log.Println(message)
		fmt.Fprint(w, message)
		return
	}
	paths := strings.Split(req.URL.Path[1:], "/")

	if len(paths) != 2 {
		w.WriteHeader(http.StatusNotFound)
		message := "404 page not found"
		log.Println(message)
		fmt.Fprint(w, message)
		return
	}

	ID, err := strconv.ParseUint(paths[1], 10, 64)
	if err != nil {
		message := fmt.Sprintf("invalid ID: %v", err)
		log.Println(message)
		fmt.Fprint(w, message)
		return
	}

	p, err := pc.interactor.GetById(uint(ID))
	if err != nil {
		message := fmt.Sprintf("could not fetch all pokemon: %v", err)
		log.Println(message)
		fmt.Fprint(w, message)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if p.ID != uint(ID) {
		data, err := json.Marshal(nil)
		if err != nil {
			message := fmt.Sprintf("unable to marshall response: %v", err)
			log.Println(message)
			fmt.Fprint(w, message)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
		return
	}

	data, err := json.Marshal(p)
	if err != nil {
		message := fmt.Sprintf("unable to marshall response: %v", err)
		log.Println(message)
		fmt.Fprint(w, message)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (pc *pokemonController) CreatePokemon(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		message := fmt.Sprintf("unsupported method [%v]", req.Method)
		log.Println(message)
		fmt.Fprint(w, message)
		return
	}

	var p model.Pokemon
	err := json.NewDecoder(req.Body).Decode(&p)
	if err != nil {
		message := fmt.Sprintf("unable to unmarshall body: %v", err)
		log.Println(message)
		fmt.Fprint(w, message)
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = pc.interactor.Create(&p)
	if err != nil {
		message := fmt.Sprintf("unable to save pokemon data: %v", err)
		log.Println(message)
		fmt.Fprint(w, message)
		w.WriteHeader(http.StatusInternalServerError)
	}

	data, err := json.Marshal(p)
	if err != nil {
		message := fmt.Sprintf("unable to marshall response: %v", err)
		log.Println(message)
		fmt.Fprint(w, message)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
