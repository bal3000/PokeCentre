package main

import (
	"context"
	"sync"

	"github.com/bal3000/PokeCentre/proto/pokemon"
	"google.golang.org/protobuf/types/known/emptypb"
)

// TODO add db model and call db from here
type PokemonService struct {
	mu      sync.Mutex
	pokemon []*pokemon.Pokemon
	pokemon.UnimplementedPokemonServiceServer
}

func NewPokemonService() *PokemonService {
	return &PokemonService{
		pokemon: []*pokemon.Pokemon{
			{
				Id:          1,
				Name:        "Bulbasaur",
				Number:      1,
				Type:        "Grass",
				Description: "Bulbasaur can be seen nearing the edge of a cliff",
			},
			{
				Id:          2,
				Name:        "Ivysaur",
				Number:      2,
				Type:        "Grass",
				Description: "Ivysaur is a Bulbasaur that has evolved into a form that can grow up to 2 meters tall",
			},
			{
				Id:          3,
				Name:        "Venusaur",
				Number:      3,
				Type:        "Grass",
				Description: "Venusaur is a Bulbasaur that has evolved into a form that can grow up to 2 meters tall",
			},
		},
	}
}

func (p *PokemonService) GetAllPokemon(context context.Context, _ *emptypb.Empty) (*pokemon.PokemonList, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	return &pokemon.PokemonList{
		Pokemon: p.pokemon,
	}, nil
}

func (p *PokemonService) GetPokemon(context context.Context, request *pokemon.GetPokemonRequest) (*pokemon.Pokemon, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	var pokemon *pokemon.Pokemon
	for _, p := range p.pokemon {
		if p.Number == request.Number {
			pokemon = p
			break
		}
	}

	return pokemon, nil
}

func (p *PokemonService) GetPokemonByType(context context.Context, request *pokemon.GetPokemonByTypeRequest) (*pokemon.PokemonList, error) {
	pokemons := []*pokemon.Pokemon{}

	for _, p := range p.pokemon {
		hasAllTypes := true
		for _, t := range request.Types {
			if p.Type != t {
				hasAllTypes = false
				break
			}
		}

		if hasAllTypes {
			pokemons = append(pokemons, p)
		}
	}

	return &pokemon.PokemonList{Pokemon: pokemons}, nil
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
