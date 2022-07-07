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

func paramErrorMessage(param string) error {
	return fmt.Errorf("Please provide a valid %s", "number")
}

func (h *handler) GetPokemon(c *gin.Context) {
	number := c.Param("number")
	if number == "" {
		c.AbortWithError(http.StatusBadRequest, paramErrorMessage("number"))
		return
	}

	conv, err := strconv.Atoi(number)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, paramErrorMessage("number"))
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 1*time.Second)
	defer cancel()

	p, err := h.pokemonClient.GetPokemon(ctx, &pokemon.GetPokemonRequest{Number: int32(conv)})
	if err != nil {
		fmt.Println("error occurred getting pokemon details:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(200, p)
}
