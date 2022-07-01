package main

import (
	"github.com/bal3000/PokeCentre/services/pokemon/proto/v1/pokemon"
	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()
	pokemonService := NewPokemonService()

	pokemon.RegisterPokemonServiceServer(server, pokemonService)
}
