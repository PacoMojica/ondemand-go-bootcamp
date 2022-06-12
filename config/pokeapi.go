package config

import (
	"fmt"
	"go-bootcamp/usecase/clients"
	"log"

	"github.com/spf13/viper"
)

var PokeAPI clients.PokeAPIConfig

func readPokeAPI() {
	p := &PokeAPI

	pa := viper.New()
	pa.SetConfigName("pokeapi")
	pa.SetConfigType("yml")
	pa.AddConfigPath(".")
	pa.AutomaticEnv()

	if err := pa.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := pa.Unmarshal(&p); err != nil {
		log.Fatalln(err)
	}
}
