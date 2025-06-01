package repository

import (
	"github.com/hoyirul/go-starter-kit/config"
	"github.com/hoyirul/go-starter-kit/internal/models"
	"github.com/hoyirul/go-starter-kit/utils"

	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	FindAll(c *gin.Context, search string) ([]models.User, *utils.Pagination, error)
	FindByID(id string) (*models.User, error)
	Create(user *models.User) error
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindAll(c *gin.Context, search string) ([]models.User, *utils.Pagination, error) {
	var users []models.User
	query := config.DB.Model(&models.User{})

	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	pagination, err := utils.Paginate(c, query, &models.User{}, &users)
	if err != nil {
		return nil, nil, err
	}

	return users, pagination, nil
}

func (r *userRepository) FindByID(id string) (*models.User, error) {
	var user models.User
	if err := config.DB.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user *models.User) error {
	return config.DB.Create(user).Error
}
