package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/bal3000/PokeCentre/api/poke-centre/models"
	"github.com/bal3000/PokeCentre/proto/trainers"
	"github.com/gin-gonic/gin"
)

func (h *handler) AddTrainer(c *gin.Context) {
	var addTrainer models.TrainerAddModel

	err := c.BindJSON(&addTrainer)
	if err != nil {
		Validator.AddErrorMessage("Trainer")
		c.AbortWithStatusJSON(http.StatusBadRequest, Validator)
		return
	}

	if addTrainer.Name == "" {
		Validator.AddErrorMessage("name")
	}
	if addTrainer.Email == "" {
		Validator.AddErrorMessage("email")
	}
	if addTrainer.Address == "" {
		Validator.AddErrorMessage("address")
	}
	if addTrainer.Phone == "" {
		Validator.AddErrorMessage("phone")
	}
	if addTrainer.NhsNumber == "" {
		Validator.AddErrorMessage("nhs number")
	}

	if len(Validator.Errors) > 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, Validator)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 1*time.Second)
	defer cancel()

	res, err := h.trainersClient.AddTrainer(ctx, &trainers.AddTrainerRequest{
		Name:      addTrainer.Name,
		Email:     addTrainer.Email,
		Address:   addTrainer.Address,
		Phone:     addTrainer.Phone,
		NhsNumber: addTrainer.NhsNumber,
	})
	if err != nil {
		fmt.Println("error occurred addding trainer:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(201, res)
}
