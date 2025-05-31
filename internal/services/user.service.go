package services

import (
	"go-starter-kit/internal/models"
	"go-starter-kit/internal/repository"
	"go-starter-kit/utils"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

type UserService interface {
	GetUsers(c *gin.Context, search string) ([]models.User, *utils.Pagination, error)
	GetUser(id string) (*models.User, error)
	CreateUser(user *models.User) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) GetUsers(c *gin.Context, search string) ([]models.User, *utils.Pagination, error) {
	return s.repo.FindAll(c, search)
}

func (s *userService) GetUser(id string) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) CreateUser(user *models.User) error {
	user.ID = uuid.NewString()
	return s.repo.Create(user)
}
