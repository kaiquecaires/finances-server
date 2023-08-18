package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kaiquecaires/finances-server/helpers"
	"github.com/kaiquecaires/finances-server/models"
	"github.com/kaiquecaires/finances-server/repositories"
)

type ListTransactionsHandler struct {
	TransactionsRepository repositories.TransactionsRepository
}

func (d *ListTransactionsHandler) Handler(c *gin.Context) {
	userId := c.MustGet("UserID").(string)
	var queryParams models.ListTransactionsModel

	if err := c.ShouldBindQuery(&queryParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON data"})
	}

	validate := validator.New()

	if err := validate.Struct(queryParams); err != nil {
		errs := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"error": helpers.ValidationErrorsToString(errs)})
		return
	}

	transactions, err := d.TransactionsRepository.List(userId, queryParams.Limit, queryParams.Page)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching transactions"})
		return
	}

	totalAmount, err := d.TransactionsRepository.GetAmount(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching transactions"})
		return
	}

	total, err := d.TransactionsRepository.GetTotal(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching transactions"})
		return
	}

	pages := total / queryParams.Limit

	if pages == 0 {
		pages = 1
	}

	c.JSON(http.StatusOK, gin.H{
		"totalAmount":  totalAmount,
		"pages":        pages,
		"page":         queryParams.Page,
		"transactions": transactions,
		"total":        total,
	})
}
