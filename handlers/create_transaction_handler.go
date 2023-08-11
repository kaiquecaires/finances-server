package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kaiquecaires/finances-server/helpers"
	"github.com/kaiquecaires/finances-server/models"
	"github.com/kaiquecaires/finances-server/repositories"
)

type CreateTransactionHandler struct {
	TransactionsRepository repositories.TransactionsRepository
}

func (d *CreateTransactionHandler) Handler(c *gin.Context) {
	var data models.CreateTransactionModel

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON data"})
		return
	}

	validate := validator.New()

	if err := validate.Struct(data); err != nil {
		errs := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"error": helpers.ValidationErrorsToString(errs)})
		return
	}

	userID := c.MustGet("UserID").(string)
	data.UserId = userID

	createdData, err := d.TransactionsRepository.Create(data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error to create the transaction"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"transaction": createdData})
}
