package routes

import (
	handler "go-starter-kit/internal/handlers"
	"go-starter-kit/internal/repository"
	service "go-starter-kit/internal/services"

	"go-starter-kit/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(rg *gin.RouterGroup) {
	authRepo := repository.NewAuthRepository()
	authService := service.NewAuthService(authRepo)
	authHandler := handler.NewAuthHandler(authService)


	auth := rg.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/logout", middlewares.JWTAuthMiddleware(authService), authHandler.Logout)
	}
}
