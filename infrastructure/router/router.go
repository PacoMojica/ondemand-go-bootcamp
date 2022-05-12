package router

import (
	"fmt"
	"go-bootcamp/config"
	"go-bootcamp/interface/controller"
	"net/http"
	"time"
)

func Init(c controller.AppController) (server *http.Server, address string) {
	mux := &http.ServeMux{}
	mux.HandleFunc("/pokemon", c.Pokemon.GetPokemon)
	mux.HandleFunc("/pokemon/", c.Pokemon.GetPokemonById)
	mux.HandleFunc("/create-pokemon", c.Pokemon.CreatePokemon)

	handler := logHandler(mux)

	port := config.Config.Server.Port
	host := config.Config.Server.Host

	address = fmt.Sprintf("%v:%v", host, port)
	server = &http.Server{
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 120 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      handler,
		Addr:         address,
	}

	return
}
