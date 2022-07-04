package main

import (
	"fmt"
	"net"
	"os"

	"github.com/bal3000/PokeCentre/proto/pokemon"
	"google.golang.org/grpc"
)

var port = ":8080"

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Using default port:", port)
	} else {
		port = os.Args[1]
	}

	server := grpc.NewServer()
	pokemonService := NewPokemonService()

	pokemon.RegisterPokemonServiceServer(server, pokemonService)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Listening on port:", port)
	server.Serve(listener)
}
