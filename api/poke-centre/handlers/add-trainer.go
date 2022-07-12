package handlers

import (
	"context"
	"time"

	"github.com/bal3000/PokeCentre/proto/trainers"
	"github.com/gin-gonic/gin"
)

func (h *handler) AddTrainer(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 1*time.Second)
	defer cancel()

	h.trainersClient.AddTrainer(ctx, &trainers.AddTrainerRequest{})
}
