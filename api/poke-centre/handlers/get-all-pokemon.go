package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *handler) GetAllPokemon(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
	defer cancel()

	maxSizeOption := grpc.MaxCallRecvMsgSize(32 * 10e6)
	list, err := h.pokemonClient.GetAllPokemon(ctx, &emptypb.Empty{}, maxSizeOption)
	if err != nil {
		fmt.Println("error occurred getting pokemon list:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(200, list.Pokemon)
}
