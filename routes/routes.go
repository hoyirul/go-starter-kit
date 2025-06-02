package routes

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hoyirul/go-starter-kit/pkg/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Set up CORS middleware
	r.Use(middlewares.CORSMiddleware())

	apiVersion := os.Getenv("API_VERSION")
	if apiVersion == "" {
		apiVersion = "v1"
	}

	// add default route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Welcome to the API version %s", apiVersion),
		})
	})

	api := r.Group(fmt.Sprintf("/api/%s", apiVersion))

	// Auth routes
	AuthRoutes(api)
	// User routes
	UserRoutes(api)
	// Product routes
	ProductRoutes(api)

	return r
}
