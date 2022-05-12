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
	re := registry.New(db)
	c := re.NewAppController()
	ro := router.New(c)
	server, address := ro.Init()

	log.Printf("running server at %v", address)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
