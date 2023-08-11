package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kaiquecaires/finances-server/db"
	"github.com/kaiquecaires/finances-server/middlewares"
	"github.com/kaiquecaires/finances-server/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	pool, err := db.CreatePostgresPool()
	if err != nil {
		log.Fatalf("Error to create pool: %v", pool)
	}

	r := gin.Default()
	r.Use(middlewares.SetDBPool)
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API IS ON FIRE",
		})
	})

	routes.SetupUserRoutes(r, pool)
	routes.SetupTransactionsRoutes(r, pool)
	r.Run(os.Getenv("APP_PORT"))
}
