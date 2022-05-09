package config

import (
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

type config struct {
	Database struct {
		File string
	}
	Server struct {
		Port string
		Host string
	}
}

var Config config

func Read() {
	c := &Config

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&c); err != nil {
		log.Fatalln(err)
	}

	spew.Dump(Config)
}
