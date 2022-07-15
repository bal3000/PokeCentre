package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/bal3000/PokeCentre/api/poke-centre/models"
	"github.com/bal3000/PokeCentre/proto/trainers"
	"github.com/gin-gonic/gin"
)

func (h *handler) UpdateTrainer(c *gin.Context) {
	id := c.Param("id")
	conv, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		Validator.AddErrorMessage("id")
		c.AbortWithStatusJSON(http.StatusBadRequest, Validator)
		return
	}

	var updateTrainer models.TrainerUpdateModel

	err = c.BindJSON(&updateTrainer)
	if err != nil {
		Validator.AddErrorMessage("Trainer")
		c.AbortWithStatusJSON(http.StatusBadRequest, Validator)
		return
	}

	if !validateTrainer(
		updateTrainer.Name,
		updateTrainer.Email,
		updateTrainer.Address,
		updateTrainer.Phone,
		updateTrainer.NhsNumber,
	) {
		c.AbortWithStatusJSON(http.StatusBadRequest, Validator)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 1*time.Second)
	defer cancel()

	res, err := h.trainersClient.UpdateTrainer(ctx, &trainers.UpdateTrainerRequest{
		Id:        conv,
		Name:      updateTrainer.Name,
		Email:     updateTrainer.Email,
		Address:   updateTrainer.Address,
		Phone:     updateTrainer.Phone,
		NhsNumber: updateTrainer.NhsNumber,
	})
	if err != nil {
		fmt.Printf("Error occurred updating trainer %d: %v\n", conv, err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusNoContent, res)
}
