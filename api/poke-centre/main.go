package main

import (
	"fmt"
	"net/http"

	"github.com/bal3000/PokeCentre/api/poke-centre/handlers"
	"github.com/bal3000/PokeCentre/proto/pokemon"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"google.golang.org/grpc"
)

func createPokemonClient(port string) (pokemon.PokemonServiceClient, func(), error) {
	pokemonConn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	return pokemon.NewPokemonServiceClient(pokemonConn), func() {
		pokemonConn.Close()
	}, nil
}

func main() {
	router := gin.Default()
	router.Use(gin.CustomRecovery(func(c *gin.Context, err any) {
		// close db connections
		// log out error
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	// pokemon client
	pokemonClient, pokemonCloser, err := createPokemonClient(":8080")
	if err != nil {
		fmt.Println("problem creating pokemon client:", err)
		return
	}
	defer pokemonCloser()

	// redis
	opt, err := redis.ParseURL("redis://default:redispw@localhost:49153")
	if err != nil {
		fmt.Println("problem parsing redis connection string:", err)
		return
	}

	rdb := redis.NewClient(opt)

	// route handler for all requests
	handler := handlers.NewHandler(pokemonClient, rdb)

	p := router.Group("/pokemon")
	{
		p.GET("/", handler.GetAllPokemon)
		p.GET("/:number", handler.GetPokemon)
		p.POST("/search", handler.GetPokemonByType)
	}

	router.Run(":3000")
}
