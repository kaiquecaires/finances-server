package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kaiquecaires/finances-server/helpers"
	"github.com/kaiquecaires/finances-server/models"
	"github.com/kaiquecaires/finances-server/repositories"
)

type DeleteTransactionHandler struct {
	TransactionsRepository repositories.TransactionsRepository
}

func (d *DeleteTransactionHandler) Handle(c *gin.Context) {
	var data models.DeleteTransactionModel

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON data"})
	}

	validate := validator.New()

	if err := validate.Struct(data); err != nil {
		errs := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"error": helpers.ValidationErrorsToString(errs)})
		return
	}

	err := d.TransactionsRepository.Delete(data.Id)

	fmt.Println(err)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error to delete transaction"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
