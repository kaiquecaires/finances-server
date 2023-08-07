package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kaiquecaires/finances-server/helpers"
	"github.com/kaiquecaires/finances-server/models"
	"github.com/kaiquecaires/finances-server/repositories"
)

type LoginHandler struct {
	UsersRepository *repositories.UsersRepository
}

func (d *LoginHandler) Handler(c *gin.Context) {
	var data models.LoginModel

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON data"})
	}

	validate := validator.New()

	if err := validate.Struct(data); err != nil {
		errs := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"error": helpers.ValidationErrorsToString(errs)})
		return
	}

	userWithPassword, err := d.UsersRepository.GetWithPassword(data.Email)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error to find user:"})
		return
	}

	isPasswordCorrect := helpers.CheckPasswordHash(data.Password, userWithPassword.PasswordHash)

	if !isPasswordCorrect {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	s, err := helpers.GenerateSignedJWT(jwt.MapClaims{
		"user_id": userWithPassword.User.Id,
	})

	fmt.Println(err)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": s})
}
