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

func createDbUrl() string {
	user := os.Getenv("PG_USER")
	if user == "" {
		log.Fatalln("Environment variable PG_USER is missing")
	}

	host := os.Getenv("PG_HOST")
	if host == "" {
		log.Fatalln("Environment variable PG_HOST is missing")
	}

	pgPort := os.Getenv("PG_PORT")
	if pgPort == "" {
		log.Fatalln("Environment variable PG_PORT is missing")
	}

	database := os.Getenv("PG_DATABASE")
	if database == "" {
		log.Fatalln("Environment variable PG_DATABASE is missing")
	}

	password := os.Getenv("PG_PASSWORD")
	if password == "" {
		log.Fatalln("Environment variable PG_PASSWORD is missing")
	}

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		pgPort,
		user,
		password,
		database,
	)
}

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

	dbUrl := createDbUrl()
	redisUrl := createRedisUrl()

	// db
	db, err := data.OpenDB(data.DbConfig{
		DSN:          dbUrl,
		MaxOpenConns: 25,
		MaxIdleConns: 25,
		MaxIdleTime:  "15m",
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	log.Println("database connection pool established")

	// redis
	rdb, err := common.CreateRedisClient(redisUrl)
	if err != nil {
		fmt.Println("problem parsing redis connection string:", err)
		return
	}

	model := data.NewTrainersModel(db, rdb)

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
