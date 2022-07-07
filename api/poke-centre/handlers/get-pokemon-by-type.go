package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/bal3000/PokeCentre/proto/pokemon"
	"github.com/gin-gonic/gin"
)

func (h *handler) GetPokemonByType(c *gin.Context) {
	var search struct {
		Types []string `json:"types"`
	}

	err := c.BindJSON(&search)
	if err != nil {
		Validator.AddErrorMessage("Types")
		c.AbortWithStatusJSON(http.StatusBadRequest, Validator)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 4*time.Second)
	defer cancel()

	list, err := h.pokemonClient.GetPokemonByType(ctx, &pokemon.GetPokemonByTypeRequest{Types: search.Types})
	if err != nil {
		fmt.Println("error occurred getting pokemon list by type:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(200, list.Pokemon)
}
