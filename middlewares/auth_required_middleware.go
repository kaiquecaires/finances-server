package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kaiquecaires/finances-server/helpers"
)

func AuthRequired(c *gin.Context) {
	header := c.GetHeader("Authorization")
	headerValues := strings.Split(header, " ")

	if len(headerValues) < 2 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You must to provided a valid token"})
		c.Abort()
		return
	}

	claims, err := helpers.ValidateJWT(headerValues[1])

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You must to provided a valid token"})
		c.Abort()
		return
	}

	userID, ok := (*claims)["user_id"].(string)

	if ok {
		c.Set("UserID", userID)
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error to parse the userID"})
		c.Abort()
	}
}
