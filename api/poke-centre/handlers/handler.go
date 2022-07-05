package handlers

import (
	"github.com/bal3000/PokeCentre/proto/pokemon"
	"github.com/go-redis/redis/v9"
)

// Add redis to this and the other grpc clients

type handler struct {
	pokemonClient pokemon.PokemonServiceClient
	redisClient   *redis.Client
}

func NewHandler(pokemonClient pokemon.PokemonServiceClient, redis *redis.Client) *handler {
	return &handler{pokemonClient: pokemonClient, redisClient: redis}
}
