package handlers

import (
	"net/http"

	"github.com/hoyirul/go-starter-kit/internal/models"
	service "github.com/hoyirul/go-starter-kit/internal/services"
	"github.com/hoyirul/go-starter-kit/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	search := c.Query("search")
	users, pagination, err := h.Service.GetUsers(c, search)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to fetch users")
		return
	}

	if len(users) == 0 {
		utils.RespondWithSuccess(c, http.StatusOK, "No users found", gin.H{
			"data":       users,
			"pagination": pagination,
		})
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Users fetched successfully", gin.H{
		"data":       users,
		"pagination": pagination,
	})
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.Service.GetUser(id)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "User not found")
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "User fetched successfully", user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := h.Service.CreateUser(&user); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create user")
		return
	}

	utils.RespondWithSuccess(c, http.StatusCreated, "User created successfully", user)
}
