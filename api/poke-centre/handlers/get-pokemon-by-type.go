package handlers

import "github.com/gin-gonic/gin"

func (app *handler) GetPokemonByType(c *gin.Context) {
	var search struct {
		Types []string
	}

	c.BindJSON(&search)
	c.JSON(200, search)
}
