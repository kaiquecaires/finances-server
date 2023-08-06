package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaiquecaires/finances-server/routes"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API IS ON FIRE",
		})
	})
	routes.UserRoutes(r)
	r.Run("localhost:3000")
}
