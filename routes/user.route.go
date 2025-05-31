package routes

import (
	handler "go-starter-kit/internal/handlers"
	"go-starter-kit/internal/repository"
	service "go-starter-kit/internal/services"
	"go-starter-kit/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoutes(rg *gin.RouterGroup) {
	authService := service.NewAuthService(repository.NewAuthRepository())
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	users := rg.Group("/users", middlewares.JWTAuthMiddleware(authService))
	{
		users.GET("/", userHandler.GetUsers)
		users.GET("/:id", userHandler.GetUser)
		users.POST("/", userHandler.CreateUser)
	}
}
