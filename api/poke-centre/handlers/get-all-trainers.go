package handlers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/bal3000/PokeCentre/proto/trainers"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *handler) GetAllTrainers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 1*time.Second)
	defer cancel()

	client, err := h.trainersClient.GetAllTrainers(ctx, &emptypb.Empty{})
	if err != nil {
		fmt.Println("error occurred getting trainers list:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	t := make([]*trainers.GetTrainerResponse, 0)

	for {
		msg, err := client.Recv()
		if err != nil {
			if err == io.EOF {
				c.JSON(http.StatusOK, t)
				return
			}

			fmt.Println("error occurred getting trainer:", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		t = append(t, msg)
	}
}
