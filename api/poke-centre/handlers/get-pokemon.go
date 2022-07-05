package handlers

import (
	"github.com/gin-gonic/gin"
)

func (app *handler) GetPokemon(c *gin.Context) {
	number := c.Param("number")

	c.JSON(200, number)
}
