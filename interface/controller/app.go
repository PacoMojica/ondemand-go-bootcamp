package controller

type AppController struct {
	Pokemon interface{ PokemonController }
	PokeAPI interface{ PokeAPIController }
}
