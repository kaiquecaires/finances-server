package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kaiquecaires/finances-server/helpers"
	"github.com/kaiquecaires/finances-server/models"
	"github.com/kaiquecaires/finances-server/usecases"
)

func SignupHandler(c *gin.Context) {
	var data models.SignupModel

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON data"})
	}

	validate := validator.New()

	if err := validate.Struct(data); err != nil {
		errs := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"error": helpers.ValidationErrorsToString(errs)})
		return
	}

	user := usecases.Signup(&data)
	c.JSON(http.StatusOK, gin.H{"user": user})
}
