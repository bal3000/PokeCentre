package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/bal3000/PokeCentre/proto/trainers"
	"github.com/gin-gonic/gin"
)

func (h *handler) DeleteTrainer(c *gin.Context) {
	id := c.Param("id")
	conv, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		Validator.AddErrorMessage("id")
		c.AbortWithStatusJSON(http.StatusBadRequest, Validator)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 1*time.Second)
	defer cancel()

	res, err := h.trainersClient.DeleteTrainer(ctx, &trainers.DeleteTrainerRequest{Id: conv})
	if err != nil {
		fmt.Printf("Error occurred deleting trainer %d: %v\n", conv, err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusAccepted, res)
}
