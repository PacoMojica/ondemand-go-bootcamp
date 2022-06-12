package clients

import (
	"errors"
	"fmt"
	"go-bootcamp/config"
	"go-bootcamp/usecase/clients"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type pokeAPIClient struct {
	config clients.PokeAPIConfig
}

func NewPokeAPIClient(c clients.PokeAPIConfig) clients.PokeAPIClient {
	return &pokeAPIClient{c}
}

func fetch(URL string) (*http.Response, error) {
	res, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(b))
	}

	return res, nil
}

func (c *pokeAPIClient) FetchPokemon(ID any) (*http.Response, error) {
	URL := fmt.Sprintf(
		"%v/%v/%v",
		c.config.BaseURL, c.config.Endpoints.Pokemon, ID)
	log.Println(fmt.Sprintf("fetching PokeAPI: '%v'", URL))
	return fetch(URL)
}

func (c *pokeAPIClient) FetchRandomPokemon() (*http.Response, error) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	ID := (r.Intn(config.PokeAPI.TotalPokemon - 1)) + 1
	URL := fmt.Sprintf(
		"%v/%v/%v",
		c.config.BaseURL, c.config.Endpoints.Pokemon, ID)
	log.Println(fmt.Sprintf("fetching PokeAPI: '%v'", URL))
	return fetch(URL)
}
