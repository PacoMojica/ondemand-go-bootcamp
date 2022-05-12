package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type pokeAPI struct {
	BaseURL      string
	TotalPokemon int
	Endpoints    struct {
		Pokemon string
		Species string
	}
}

var PokeAPI pokeAPI

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
