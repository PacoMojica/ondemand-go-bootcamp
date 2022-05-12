package controller

import (
	"errors"
	"fmt"
	"go-bootcamp/config"
	"io/ioutil"
	"log"
	"net/http"
)

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

func fetchPokeAPI(endpoint string) (*http.Response, error) {
	URL := fmt.Sprintf("%v/%v", config.PokeAPI.BaseURL, endpoint)
	log.Println(fmt.Sprintf("fetching PokeAPI: '%v'", URL))
	return fetch(URL)
}
