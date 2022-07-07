package handlers

import (
	"fmt"

	"github.com/bal3000/PokeCentre/proto/pokemon"
)

var Validator *ValidationError = &ValidationError{}

type handler struct {
	pokemonClient pokemon.PokemonServiceClient
}

func NewHandler(pokemonClient pokemon.PokemonServiceClient) *handler {
	return &handler{pokemonClient: pokemonClient}
}

type ValidationError struct {
	Errors []string `json:"errors"`
}

func (v *ValidationError) AddErrorMessage(param string) {
	v.Errors = append(v.Errors, fmt.Sprintf("Please provide a valid %s", "number"))
}
