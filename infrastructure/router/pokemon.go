package router

func (ar *appRouter) handlePokemon() {
	ar.mux.HandleFunc("/pokemon", ar.controller.Pokemon.GetPokemon)
	ar.mux.HandleFunc("/pokemon/", ar.controller.Pokemon.GetPokemonById)
	ar.mux.HandleFunc("/create-pokemon", ar.controller.Pokemon.CreatePokemon)
}
