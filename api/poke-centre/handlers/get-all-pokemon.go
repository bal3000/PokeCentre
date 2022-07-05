package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *handler) GetAllPokemon(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 1*time.Second)
	defer cancel()

	val, err := h.redisClient.Get(ctx, "all-pokemon").Result()
	if err != nil {
		if err == redis.Nil {
			list, err := h.pokemonClient.GetAllPokemon(ctx, &emptypb.Empty{})
			if err != nil {
				fmt.Println("error occurred getting pokemon list:", err)
				c.AbortWithStatus(http.StatusInternalServerError)
			}

			h.redisClient.Set(ctx, "all-pokemon", list, 24*time.Hour)

			c.JSON(200, list.Pokemon)
			return
		}

		fmt.Println("error occured getting key from redis:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(200, val)
}
