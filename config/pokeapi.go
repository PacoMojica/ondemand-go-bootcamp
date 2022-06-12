package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"

	"go-bootcamp/usecase/clients"
)

var PokeAPI clients.PokeAPIConfig

func readPokeAPI() {
	p := &PokeAPI

	pa := viper.New()
	pa.SetConfigName("pokeapi")
	pa.SetConfigType("yml")
	pa.AddConfigPath("./config")
	pa.AutomaticEnv()

	if err := pa.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := pa.Unmarshal(&p); err != nil {
		log.Fatalln(err)
	}
}
