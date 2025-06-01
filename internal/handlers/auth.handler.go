package handlers

import (
	"fmt"
	"net/http"

	"github.com/hoyirul/go-starter-kit/internal/models"
	"github.com/hoyirul/go-starter-kit/internal/services"
	"github.com/hoyirul/go-starter-kit/utils"

	"strings"
	"time"

	"github.com/hoyirul/go-starter-kit/config"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthHandler struct {
	service services.AuthService
}

func NewAuthHandler(s services.AuthService) *AuthHandler {
	return &AuthHandler{service: s}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var request struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	userID := uuid.New()

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request")
		return
	}

	existingUser, err := h.service.FindByEmail(request.Email)
	fmt.Printf("Existing user: %+v, Error: %v\n", existingUser, err)
	if err == nil && existingUser != nil {
		utils.RespondWithError(c, http.StatusConflict, "Email already registered")
		return
	}

	user := models.User{
		ID:         userID.String(),
		Name:       request.Name,
		Email:      request.Email,
		Password:   request.Password,
		CreatedAt: 	time.Now(),
		UpdatedAt: 	time.Now(),
	}

	if err := h.service.Register(&user); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	response := gin.H{
		"id":         user.ID,
		"name":       user.Name,
		"email":      user.Email,
		"created_at": user.CreatedAt.Format(time.RFC3339),
		"updated_at": user.UpdatedAt.Format(time.RFC3339),
	}

	utils.RespondWithSuccess(c, http.StatusCreated, "User registered successfully", response)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request")
		return
	}

	user, err := h.service.Login(request.Email, request.Password)
	if err != nil {
		utils.RespondWithError(c, http.StatusUnauthorized, err.Error())
		return
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	response := gin.H{
		"access_token": token,
		"token_type":  "Bearer",
		"expires_in":  24 * time.Hour.Seconds(),
		"user": gin.H{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"created_at": user.CreatedAt.Format(time.RFC3339),
			"updated_at": user.UpdatedAt.Format(time.RFC3339),
		},
	}

	ctx := c.Request.Context()
	err = config.RedisClient.Set(ctx, token, "valid", 24*time.Hour).Err()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to store token in Redis")
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Login successful", response)
}

func (h *AuthHandler) GetUserProfile(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		utils.RespondWithError(c, http.StatusUnauthorized, "User not authenticated")
		return
	}

	u, ok := user.(models.User)
	if !ok {
		utils.RespondWithError(c, http.StatusUnauthorized, "Invalid user data")
		return
	}

	response := gin.H{
		"id":         u.ID,
		"name":       u.Name,
		"email":      u.Email,
		"created_at": u.CreatedAt.Format(time.RFC3339),
		"updated_at": u.UpdatedAt.Format(time.RFC3339),
	}

	utils.RespondWithSuccess(c, http.StatusOK, "User profile fetched successfully", response)
}

func (h *AuthHandler) Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid authorization header")
		return
	}

	tokenString := parts[1]
	ctx := c.Request.Context()

	err := config.RedisClient.Del(ctx, tokenString).Err()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to logout")
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Logout successful", nil)
}

