package handlers

import (
	"fmt"

	"github.com/bal3000/PokeCentre/proto/pokemon"
	"github.com/bal3000/PokeCentre/proto/trainers"
)

var Validator *ValidationError = &ValidationError{}

type handler struct {
	pokemonClient  pokemon.PokemonServiceClient
	trainersClient trainers.TrainersServiceClient
}

func NewHandler(pokemonClient pokemon.PokemonServiceClient, trainersClient trainers.TrainersServiceClient) *handler {
	return &handler{
		pokemonClient:  pokemonClient,
		trainersClient: trainersClient,
	}
}

type ValidationError struct {
	Errors []string `json:"errors"`
}

func (v *ValidationError) AddErrorMessage(param string) {
	v.Errors = append(v.Errors, fmt.Sprintf("Please provide a valid %s", param))
}

func validateTrainer(name, email, address, phone, nhsNumber string) bool {
	if name == "" {
		Validator.AddErrorMessage("name")
	}
	if email == "" {
		Validator.AddErrorMessage("email")
	}
	if address == "" {
		Validator.AddErrorMessage("address")
	}
	if phone == "" {
		Validator.AddErrorMessage("phone")
	}
	if nhsNumber == "" {
		Validator.AddErrorMessage("nhs number")
	}

	return len(Validator.Errors) == 0
}
