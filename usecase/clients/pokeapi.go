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
	FetchPokemon(identifier any) (*http.Response, error)
	FetchRandomPokemon() (*http.Response, error)
}
