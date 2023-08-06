package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kaiquecaires/finances-server/routes"
)

func main() {
	fmt.Println("Starting...")
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API IS ON FIRE",
		})
	})
	routes.UserRoutes(r)
	r.Run(os.Getenv("APP_PORT"))
	fmt.Println("API IS ON FIRE")
}
