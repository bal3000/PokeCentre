package main

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/bal3000/PokeCentre/proto/pokemon"
	"github.com/bal3000/PokeCentre/services/common/data"
	"github.com/go-redis/redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PokemonService struct {
	mu                sync.Mutex
	pokemonCollection *mongo.Collection
	redisClient       *redis.Client
	pokemon.UnimplementedPokemonServiceServer
}

func NewPokemonService(pokemonCollection *mongo.Collection, redis *redis.Client) *PokemonService {
	return &PokemonService{
		pokemonCollection: pokemonCollection,
		redisClient:       redis,
	}
}

func (p *PokemonService) GetAllPokemon(context context.Context, _ *emptypb.Empty) (*pokemon.PokemonList, error) {

	// add cache here
	pm, err := data.GetAndSetValue(context, p.redisClient, "all-pokemon", func() ([]*pokemon.PokemonSimple, error) {
		cur, err := p.pokemonCollection.Find(context, bson.D{{}}, options.Find())
		if err != nil {
			return nil, err
		}
		defer cur.Close(context)

		results := make([]*pokemon.PokemonSimple, 0)
		if err = cur.All(context, &results); err != nil {
			return nil, err
		}

		return results, nil
	})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// add pagination or streaming

	fmt.Printf("Found %d results\n", len(pm))

	return &pokemon.PokemonList{
		Pokemon: pm,
	}, nil
}

func (p *PokemonService) GetPokemon(context context.Context, request *pokemon.GetPokemonRequest) (*pokemon.Pokemon, error) {
	key := fmt.Sprintf("pokemon-details-%d", request.Number)
	pokemon, err := data.GetAndSetValue(context, p.redisClient, key, func() (*pokemon.Pokemon, error) {
		var pokemon pokemon.Pokemon
		err := p.pokemonCollection.FindOne(context, bson.D{{"id", request.Number}}).Decode(&pokemon)
		if err != nil {
			return nil, err
		}
		return &pokemon, nil
	})

	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}
	fmt.Printf("Found Pokemon: %d: %s\n", pokemon.Id, pokemon.Name)

	return pokemon, nil
}

func (p *PokemonService) GetPokemonByType(context context.Context, request *pokemon.GetPokemonByTypeRequest) (*pokemon.PokemonList, error) {
	key := fmt.Sprintf("pokemon-bytype-%s", strings.Join(request.Types, "-"))
	results, err := data.GetAndSetValue(context, p.redisClient, key, func() ([]*pokemon.PokemonSimple, error) {
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

		results := make([]*pokemon.PokemonSimple, 0)
		if err = cursor.All(context, &results); err != nil {
			return nil, err
		}

		return results, nil
	})

	if err != nil {
		return nil, err
	}

	fmt.Printf("Found %d results\n", len(results))

	return &pokemon.PokemonList{Pokemon: results}, nil
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
