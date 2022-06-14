package config

import "github.com/davecgh/go-spew/spew"

// loads the values from the config files
func Read() {
	readApp()
	readPokeAPI()

	spew.Dump(App)
	spew.Dump(PokeAPI)
}
