package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bal3000/PokeCentre/api/poke-centre/handlers"
	"github.com/bal3000/PokeCentre/proto/pokemon"
	"github.com/bal3000/PokeCentre/proto/trainers"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func createPokemonClient() (pokemon.PokemonServiceClient, func(), error) {
	psUrl := os.Getenv("POKEMON_SERVER")
	if psUrl == "" {
		return nil, nil, errors.New("Environment variable POKEMON_SERVER is missing")
	}

	fmt.Printf("Connecting to %s\n", psUrl)

	pokemonConn, err := grpc.Dial(psUrl, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	return pokemon.NewPokemonServiceClient(pokemonConn), func() {
		pokemonConn.Close()
	}, nil
}

func createTrainersClient() (trainers.TrainersServiceClient, func(), error) {
	tsUrl := os.Getenv("TRAINER_SERVER")
	if tsUrl == "" {
		return nil, nil, errors.New("Environment variable TRAINER_SERVER is missing")
	}

	fmt.Printf("Connecting to %s\n", tsUrl)

	conn, err := grpc.Dial(tsUrl, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	return trainers.NewTrainersServiceClient(conn), func() { conn.Close() }, nil
}

func main() {
	router := gin.Default()
	router.Use(gin.CustomRecovery(func(c *gin.Context, err any) {
		// close db connections
		// log out error
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	// pokemon client
	pokemonClient, pokemonCloser, err := createPokemonClient()
	if err != nil {
		log.Fatalln("problem creating pokemon client:", err)
	}
	defer pokemonCloser()

	trainersClient, trainersCloser, err := createTrainersClient()
	if err != nil {
		log.Fatalln("problem creating trainers client:", err)
	}
	defer trainersCloser()

	// route handler for all requests
	handler := handlers.NewHandler(pokemonClient, trainersClient)

	p := router.Group("/pokemon")
	{
		p.GET("/", handler.GetAllPokemon)
		p.GET("/:number", handler.GetPokemon)
		p.POST("/search", handler.GetPokemonByType)
	}

	t := router.Group("/trainers")
	{
		t.GET("/", handler.GetAllTrainers)
		t.POST("/", handler.AddTrainer)
		t.PUT("/:id", handler.UpdateTrainer)
		t.DELETE("/:id", handler.DeleteTrainer)
	}

	router.Run(":3000")
}
