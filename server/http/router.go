package http

import (
	"mywallet/controller"
	// "mywallet/internal/delivery/http/middleware"
	"mywallet/middleware"
	"mywallet/server"

	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	// Set Gin mode
	gin.SetMode(server.Cfg.GinMode)

	router := gin.Default()

	// Global middlewares
	router.Use(middleware.ErrorHandler())
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.SecurityHeadersMiddleware())

	// API routes
	api := router.Group("/api")
	{
		// Auth routes (public)
		auth := api.Group("/auth")
		{
			auth.POST("/register", controller.Register)
			auth.POST("/login", controller.Login)
		}

		// Protected routes
		authMiddleware := middleware.AuthMiddleware(server.UserUsecase)

		// User routes
		users := api.Group("/users")
		users.Use(authMiddleware)
		{
			users.GET("/profile", controller.GetProfile)
		}

		// Wallet routes
		wallets := api.Group("/wallets")
		wallets.Use(authMiddleware)
		{
			wallets.GET("/balance", controller.GetBalance)
			wallets.POST("/topup", controller.TopUp)
		}

		// Transaction routes
		transactions := api.Group("/transactions")
		transactions.Use(authMiddleware)
		{
			transactions.POST("/transfer", controller.Transfer)
			transactions.GET("/history", controller.GetHistory)
		}
	}

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	return router
}
