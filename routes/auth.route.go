package routes

import (
	handler "github.com/hoyirul/go-starter-kit/internal/handlers"
	"github.com/hoyirul/go-starter-kit/internal/repository"
	service "github.com/hoyirul/go-starter-kit/internal/services"

	"github.com/hoyirul/go-starter-kit/pkg/middlewares"

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
		auth.GET("/profile", middlewares.JWTAuthMiddleware(authService), authHandler.GetUserProfile)
		auth.POST("/logout", middlewares.JWTAuthMiddleware(authService), authHandler.Logout)
	}
}
