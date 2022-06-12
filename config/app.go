package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type appConfig struct {
	Database struct {
		Path string
	}
	Server struct {
		Port string
		Host string
	}
}

var App appConfig

func readApp() {
	a := &App

	cfg := viper.New()
	cfg.SetConfigName("config")
	cfg.SetConfigType("yml")
	cfg.AddConfigPath(".")
	cfg.AutomaticEnv()

	if err := cfg.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := cfg.Unmarshal(&a); err != nil {
		log.Fatalln(err)
	}
}
