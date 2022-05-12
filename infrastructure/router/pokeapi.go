package router

func (ar *appRouter) handlePokeAPI() {
	ar.mux.HandleFunc("/fetch-pokeapi", ar.controller.PokeAPI.GetPokemon)
	ar.mux.HandleFunc("/fetch-pokeapi/", ar.controller.PokeAPI.GetPokemonFromIdentifier)
}
