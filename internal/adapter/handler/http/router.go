package http

import (
	"stori/internal/adapter/config"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter(
	config *config.Container,
	userHandler UserHandler,
	accountHandler AccountHandler,
	transactionHandler TransactionHandler,
) (*Router, error) {
	if config.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	ginConfig := cors.DefaultConfig()
	allowedOrigins := config.HTTP.AllowedOrigins
	originsList := strings.Split(allowedOrigins, ",")
	ginConfig.AllowOrigins = originsList

	router := gin.New()
	router.Use(gin.Recovery(), cors.New(ginConfig))

	v1 := router.Group("/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("", userHandler.Register)
			users.GET("/:email", userHandler.GetUserByEmail)
		}

		accounts := v1.Group("/accounts")
		{
			accounts.POST("", accountHandler.CreateAccount)
			{
				accounts.GET("/user/:user_id", accountHandler.GetAccountByUserID)
			}
		}

		transactions := v1.Group("/transactions")
		{
			transactions.POST("", transactionHandler.CreateTransaction)
			{
				transactions.GET("/account/:account_id", transactionHandler.GetTransactionsByAccountId)
				transactions.POST("/upload-file", transactionHandler.UploadFile)
				transactions.POST("send-email-summary", transactionHandler.SendEmailSummary)
			}
		}
	}

	return &Router{
		router,
	}, nil
}

func (r *Router) Serve(address string) error {
	return r.Run(address)
}
