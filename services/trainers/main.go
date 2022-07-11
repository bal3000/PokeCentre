package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/bal3000/PokeCentre/proto/trainers"
	common "github.com/bal3000/PokeCentre/services/common/data"
	"github.com/bal3000/PokeCentre/services/trainers/data"
	"google.golang.org/grpc"
)

var port = ":8081"

func main() {
	p := os.Getenv("PORT")
	if p != "" {
		port = p
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatalln("Environment variable DATABASE_URL is missing")
	}

	redisUrl := os.Getenv("REDIS_URL")
	if redisUrl == "" {
		log.Fatalln("Environment variable REDIS_URL is missing")
	}

	// ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	// defer cancel()

	// db

	// redis
	rdb, err := common.CreateRedisClient(redisUrl)
	if err != nil {
		fmt.Println("problem parsing redis connection string:", err)
		return
	}

	model := data.NewTrainersModel(nil, rdb)

	server := grpc.NewServer()
	trainersService := NewTrainerService(model)

	trainers.RegisterTrainersServiceServer(server, trainersService)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Listening on port:", port)
	server.Serve(listener)
}
