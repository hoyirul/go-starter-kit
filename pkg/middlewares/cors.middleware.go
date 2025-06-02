package middlewares

import (
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	allowedOrigins := strings.Split(os.Getenv("CORS_ALLOW_ORIGIN"), ",")
	allowedMethods := strings.Split(os.Getenv("CORS_ALLOW_METHODS"), ",")
	allowedHeaders := strings.Split(os.Getenv("CORS_ALLOW_HEADERS"), ",")

	config := cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     allowedMethods,
		AllowHeaders:     allowedHeaders,
		AllowCredentials: true,
	}

	return cors.New(config)
}