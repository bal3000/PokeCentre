package main

import (
	"fmt"

	"github.com/bal3000/PokeCentre/api/poke-centre/handlers"
	"github.com/bal3000/PokeCentre/proto/pokemon"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func createPokemonClient(port string) (pokemon.PokemonServiceClient, error) {
	pokemonConn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return pokemon.NewPokemonServiceClient(pokemonConn), nil
}

func main() {
	router := gin.Default()

	pokemonClient, err := createPokemonClient(":8080")
	if err != nil {
		fmt.Println("problem creating pokemon client:", err)
		return
	}

	// route handler for all requests
	handler := handlers.NewHandler(pokemonClient)

	p := router.Group("/pokemon")
	{
		p.GET("/", handler.GetAllPokemon)
		p.GET("/:number", handler.GetPokemon)
		p.POST("/search", handler.GetPokemonByType)
	}

	router.Run(":3000")
}
