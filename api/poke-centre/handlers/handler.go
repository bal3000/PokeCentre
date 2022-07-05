package handlers

import (
	"github.com/bal3000/PokeCentre/proto/pokemon"
)

// Add redis to this and the other grpc clients

type handler struct {
	pokemonClient pokemon.PokemonServiceClient
}

func NewHandler(pokemonClient pokemon.PokemonServiceClient) *handler {
	return &handler{pokemonClient: pokemonClient}
}
