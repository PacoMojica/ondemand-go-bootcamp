package clients

import "net/http"

type PokeAPIConfig struct {
	BaseURL      string
	TotalPokemon int
	Endpoints    struct {
		Pokemon string
		Species string
	}
}

type PokeAPIClient interface {
	// Fetches a pokemon using the provided ID form the PokeAPI
	FetchPokemon(identifier any) (*http.Response, error)
	// Fetches a random pokemon from the PokeAPI
	FetchRandomPokemon() (*http.Response, error)
}
