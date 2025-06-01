package routes

import (
	handler "github.com/hoyirul/go-starter-kit/internal/handlers"
	"github.com/hoyirul/go-starter-kit/internal/repository"
	service "github.com/hoyirul/go-starter-kit/internal/services"
	"github.com/hoyirul/go-starter-kit/pkg/middlewares"

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
