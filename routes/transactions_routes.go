package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kaiquecaires/finances-server/handlers"
	"github.com/kaiquecaires/finances-server/middlewares"
	"github.com/kaiquecaires/finances-server/repositories"
)

func SetupTransactionsRoutes(r *gin.Engine, dbPool *pgxpool.Pool) {
	authorized := r.Group("/transactions")
	authorized.Use(middlewares.AuthRequired)

	authorized.POST("/", func(ctx *gin.Context) {
		handler := handlers.CreateTransactionHandler{
			TransactionsRepository: repositories.TransactionsRepository{
				DbPool: dbPool,
			},
		}

		handler.Handler(ctx)
	})
}
