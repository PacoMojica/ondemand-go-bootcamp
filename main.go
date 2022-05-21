package main

import (
	"go-bootcamp/config"
	"go-bootcamp/infrastructure/database"
	"go-bootcamp/infrastructure/router"
	"go-bootcamp/registry"
	"log"
)

func main() {
	config.Read()
	db := database.New(config.App.Database.Path)
	re := registry.New(db, config.PokeAPI)
	c := re.NewAppController()
	ro := router.New(c)
	server, address := ro.Init(
		config.App.Server.Host,
		config.App.Server.Port)

	log.Printf("running server at %v", address)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
