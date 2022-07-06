package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/bal3000/PokeCentre/proto/pokemon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/types/known/emptypb"
)

// TODO add db model and call db from here
type PokemonService struct {
	mu                sync.Mutex
	pokemonCollection *mongo.Collection
	pokemon.UnimplementedPokemonServiceServer
}

func NewPokemonService(pokemonCollection *mongo.Collection) *PokemonService {
	return &PokemonService{
		pokemonCollection: pokemonCollection,
	}
}

func (p *PokemonService) GetAllPokemon(context context.Context, _ *emptypb.Empty) (*pokemon.PokemonList, error) {
	cur, err := p.pokemonCollection.Find(context, bson.D{{}}, options.Find())
	if err != nil {
		return nil, err
	}
	defer cur.Close(context)

	results := make([]*pokemon.Pokemon, 0)
	if err = cur.All(context, &results); err != nil {
		return nil, err
	}

	fmt.Printf("Found %d results\n", len(results))

	return &pokemon.PokemonList{
		Pokemon: results,
	}, nil
}

func (p *PokemonService) GetPokemon(context context.Context, request *pokemon.GetPokemonRequest) (*pokemon.Pokemon, error) {
	var pokemon *pokemon.Pokemon
	err := p.pokemonCollection.FindOne(context, bson.D{{"id", request.Number}}).Decode(pokemon)
	if err != nil {
		return nil, err
	}

	return pokemon, nil
}

func (p *PokemonService) GetPokemonByType(context context.Context, request *pokemon.GetPokemonByTypeRequest) (*pokemon.PokemonList, error) {
	pokemons := []*pokemon.Pokemon{}

	searchArr := make(bson.A, 0)
	for _, t := range request.Types {
		searchArr = append(searchArr, t)
	}

	cursor, err := p.pokemonCollection.Find(
		context,
		bson.D{
			{"types.name", bson.D{{"$all", searchArr}}},
		},
	)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context)

	results := make([]*pokemon.Pokemon, 0)
	if err = cursor.All(context, &results); err != nil {
		return nil, err
	}
	fmt.Printf("Found %d results\n", len(results))

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
