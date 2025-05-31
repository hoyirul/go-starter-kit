package middlewares

import (
	"go-starter-kit/config"
	"go-starter-kit/utils"
	"go-starter-kit/internal/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func JWTAuthMiddleware(authService services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.RespondWithError(c, http.StatusUnauthorized, "Authorization header is required")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.RespondWithError(c, http.StatusUnauthorized, "Invalid authorization format")
			c.Abort()
			return
		}

		tokenString := parts[1]
		ctx := c.Request.Context()
		val, err := config.RedisClient.Get(ctx, tokenString).Result()
		if err == redis.Nil {
			utils.RespondWithError(c, http.StatusUnauthorized, "Token not found or expired")
			c.Abort()
			return
		} else if err != nil {
			utils.RespondWithError(c, http.StatusInternalServerError, "Redis error: "+err.Error())
			c.Abort()
			return
		}

		if val != "valid" {
			utils.RespondWithError(c, http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		userID, err := utils.ParseToken(tokenString)
		if err != nil {
			utils.RespondWithError(c, http.StatusUnauthorized, "Invalid token: "+err.Error())
			c.Abort()
			return
		}

		user, err := authService.FindByID(userID)
		if err != nil {
			utils.RespondWithError(c, http.StatusUnauthorized, "User not found: "+err.Error())
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
