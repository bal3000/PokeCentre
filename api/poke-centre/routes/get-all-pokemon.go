package routes

import "github.com/gin-gonic/gin"

func GetAllPokemon(c *gin.Context) {
	c.JSON(200, "test")
}
