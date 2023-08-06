package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kaiquecaires/finances-server/handlers"
	"github.com/kaiquecaires/finances-server/repositories"
)

func SetupUserRoutes(r *gin.Engine, dbPool *pgxpool.Pool) {
	r.POST("/signup", func(ctx *gin.Context) {
		repository := &repositories.UserRepository{DbPool: dbPool}
		handler := &handlers.SignupHandler{
			UserRepository: repository,
		}
		handler.Handler(ctx)
	})
}
