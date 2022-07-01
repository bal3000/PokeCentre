package main

import (
	"context"
	"sync"

	"github.com/bal3000/PokeCentre/services/pokemon/proto/v1/pokemon"
	"google.golang.org/protobuf/types/known/emptypb"
)

// TODO add db model and call db from here
type PokemonService struct {
	mu      sync.Mutex
	pokemon []pokemon.Pokemon
	pokemon.UnimplementedPokemonServiceServer
}

func NewPokemonService() *PokemonService {
	return &PokemonService{}
}

func (p *PokemonService) GetAllPokemon(context.Context, *emptypb.Empty) (*pokemon.PokemonList, error) {
	return nil, nil
}

func (p *PokemonService) GetPokemon(context.Context, *pokemon.GetPokemonRequest) (*pokemon.Pokemon, error) {
	return nil, nil
}

func (p *PokemonService) GetPokemonByType(context.Context, *pokemon.GetPokemonByTypeRequest) (*pokemon.PokemonList, error) {
	return nil, nil
}

func (p *PokemonService) AddPokemon(context.Context, *pokemon.AddPokemonRequest) (*pokemon.Pokemon, error) {
	return nil, nil
}

func (p *PokemonService) UpdatePokemon(context.Context, *pokemon.Pokemon) (*emptypb.Empty, error) {
	return nil, nil
}

func (p *PokemonService) DeletePokemon(context.Context, *pokemon.DeletePokemonRequest) (*emptypb.Empty, error) {
	return nil, nil
}
