package config

import (
	"github.com/davecgh/go-spew/spew"
)

func Read() {
	readApp()
	readPokeAPI()

	spew.Dump(App)
	spew.Dump(PokeAPI)
}
