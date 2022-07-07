package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/bal3000/PokeCentre/proto/pokemon"
	"github.com/gin-gonic/gin"
)

func (h *handler) GetPokemon(c *gin.Context) {
	number := c.Param("number")

	conv, err := strconv.Atoi(number)
	if err != nil {
		Validator.AddErrorMessage("number")
		c.AbortWithStatusJSON(http.StatusBadRequest, Validator)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 1*time.Second)
	defer cancel()

	p, err := h.pokemonClient.GetPokemon(ctx, &pokemon.GetPokemonRequest{Number: int32(conv)})
	if err != nil {
		fmt.Println("error occurred getting pokemon details:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(200, p)
}
