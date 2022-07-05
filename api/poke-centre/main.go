package main

import (
	"github.com/bal3000/PokeCentre/api/poke-centre/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	p := router.Group("/pokemon")
	{
		p.GET("/", routes.GetAllPokemon)
		p.GET("/:number", routes.GetPokemon)
		p.POST("/search", routes.GetPokemonByType)
	}

	router.Run(":3000")
}
