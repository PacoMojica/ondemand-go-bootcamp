package router

import (
	"fmt"
	"go-bootcamp/config"
	"go-bootcamp/interface/controller"
	"net/http"
	"time"
)

type appRouter struct {
	mux        *http.ServeMux
	controller controller.AppController
}

type AppRouter interface {
	Init() (*http.Server, string)
	handlePokemon()
	handlePokeAPI()
}

func New(c controller.AppController) AppRouter {
	m := &http.ServeMux{}
	return &appRouter{m, c}
}

func (ar *appRouter) Init() (server *http.Server, address string) {
	ar.handlePokemon()
	ar.handlePokeAPI()

	handler := logHandler(ar.mux)

	port := config.App.Server.Port
	host := config.App.Server.Host

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
