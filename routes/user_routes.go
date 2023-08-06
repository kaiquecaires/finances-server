package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaiquecaires/finances-server/handlers"
)

func UserRoutes(r *gin.Engine) {
	r.POST("/signup", handlers.SignupHandler)
}
