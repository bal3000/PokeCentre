package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/bal3000/PokeCentre/proto/pokemon"
	"github.com/bal3000/PokeCentre/services/common/data"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var port = ":8080"

func createRedisUrl() string {
	redisUrl := os.Getenv("REDIS_HOST")
	if redisUrl == "" {
		log.Fatalln("Environment variable REDIS_HOST is missing")
	}

	redisPort := os.Getenv("REDIS_PORT")
	if redisUrl == "" {
		log.Fatalln("Environment variable REDIS_PORT is missing")
	}

	return fmt.Sprintf("redis://%s:%s", redisUrl, redisPort)
}

func main() {
	p := os.Getenv("PORT")
	if p != "" {
		port = p
	}

	mongoUri := os.Getenv("MONGODB_URI")
	if mongoUri == "" {
		log.Fatalln("Environment variable MONGODB_URI is missing")
	}

	redisUrl := createRedisUrl()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// mongo
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Fatalln("Failed to connect to mongo db:", err)
	}

	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := client.Disconnect(ctx); err != nil {
			log.Fatalln("Failed to disconnect from mongo db:", err)
		}
	}()

	col := client.Database("pokedex").Collection("pokemon")

	// redis
	rdb, err := data.CreateRedisClient(redisUrl)
	if err != nil {
		fmt.Println("problem parsing redis connection string:", err)
		return
	}

	server := grpc.NewServer()
	pokemonService := NewPokemonService(col, rdb)

	pokemon.RegisterPokemonServiceServer(server, pokemonService)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Listening on port:", port)
	server.Serve(listener)
}
