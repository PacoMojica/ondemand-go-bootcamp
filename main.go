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
	db := database.New()
	r := registry.New(db)
	c := r.NewAppController()
	server, address := router.Init(c)

	log.Printf("running server at %v", address)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
