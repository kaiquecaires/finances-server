package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kaiquecaires/finances-server/handlers"
	"github.com/kaiquecaires/finances-server/repositories"
)

func SetupUserRoutes(r *gin.Engine, dbPool *pgxpool.Pool) {
	usersRepository := &repositories.UsersRepository{DbPool: dbPool}
	userPasswordRepository := &repositories.UserPasswordsRepository{DbPool: dbPool}

	r.POST("/signup", func(ctx *gin.Context) {
		handler := &handlers.SignupHandler{
			UsersRepository:         usersRepository,
			UserPasswordsRepository: userPasswordRepository,
		}
		handler.Handler(ctx)
	})

	r.POST("/login", func(ctx *gin.Context) {
		handler := &handlers.LoginHandler{
			UsersRepository: usersRepository,
		}
		handler.Handler(ctx)
	})
}
