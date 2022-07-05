package routes

import (
	"github.com/gin-gonic/gin"
)

func GetPokemon(c *gin.Context) {
	number := c.Param("number")

	c.JSON(200, number)
}
