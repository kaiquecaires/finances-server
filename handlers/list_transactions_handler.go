package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaiquecaires/finances-server/repositories"
)

type ListTransactionsHandler struct {
	TransactionsRepository repositories.TransactionsRepository
}

func (d *ListTransactionsHandler) Handler(c *gin.Context) {
	userId := c.MustGet("UserID").(string)
	transactions, err := d.TransactionsRepository.List(userId)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error to get transactions"})
		return
	}

	totalAmount, err := d.TransactionsRepository.GetAmount(userId)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error to get transactions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactions": transactions, "totalAmount": totalAmount})
}
