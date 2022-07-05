package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/bal3000/PokeCentre/api/poke-centre/data"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *handler) GetAllPokemon(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 1*time.Second)
	defer cancel()

	val, err := data.GetOrSetValue(ctx, h.redisClient, "all-pokemon", func() (data.PokemonCollection, error) {
		list, err := h.pokemonClient.GetAllPokemon(ctx, &emptypb.Empty{})
		if err != nil {
			fmt.Println("error occurred getting pokemon list:", err)
			return data.PokemonCollection{}, err
		}

		mapped := data.MapCollectionToModel(list.Pokemon)

		return data.PokemonCollection{Pokemon: mapped}, nil
	})

	if err != nil {
		fmt.Println("redis error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(200, val)
}
