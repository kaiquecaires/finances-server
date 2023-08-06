package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kaiquecaires/finances-server/db"
)

func SetDBPool(c *gin.Context) {
	dbPool, err := db.CreatePostgresPool()

	if err != nil {
		log.Fatalf("Error connecting to db: %s", err)
	}

	c.Set("dbPool", dbPool)
	c.Next()
}
