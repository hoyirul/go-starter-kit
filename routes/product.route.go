package routes

import (
	handler "go-starter-kit/internal/handlers"
	"go-starter-kit/internal/repository"
	service "go-starter-kit/internal/services"
	"go-starter-kit/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(rg *gin.RouterGroup) {
	authService := service.NewAuthService(repository.NewAuthRepository())
	productRepo := repository.NewProductRepository()
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	products := rg.Group("/products", middlewares.JWTAuthMiddleware(authService))
	{
		products.GET("/", productHandler.GetProducts)
		products.GET("/:id", productHandler.GetProduct)
		products.POST("/", productHandler.CreateProduct)
		products.PUT("/:id", productHandler.UpdateProduct)
		products.DELETE("/:id", productHandler.DeleteProduct)
	}
}
