package repository

import (
	"go-starter-kit/config"
	"go-starter-kit/internal/models"
	"go-starter-kit/utils"

	"github.com/gin-gonic/gin"
)

type ProductRepository interface {
	FindAll(c *gin.Context, search string) ([]models.ProductResponse, *utils.Pagination, error)
	FindByID(id uint) (*models.ProductResponse, error)
	Create(product *models.Product) error
	Update(product *models.Product) error
	Delete(id uint) error
}

type productRepository struct{}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}

func (r *productRepository) FindAll(c *gin.Context, search string) ([]models.ProductResponse, *utils.Pagination, error) {
	var products []models.ProductResponse
	query := config.DB.Preload("User").Model(&models.ProductResponse{})

	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	pagination, err := utils.Paginate(c, query, &models.ProductResponse{}, &products)
	if err != nil {
		return nil, nil, err
	}

	return products, pagination, nil
}

func (r *productRepository) FindByID(id uint) (*models.ProductResponse, error) {
	var product models.ProductResponse
	if err := config.DB.Preload("User").First(&product, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Create(product *models.Product) error {
	return config.DB.Create(product).Error
}

func (r *productRepository) Update(product *models.Product) error {
	return config.DB.Model(&models.Product{}).Where("id = ?", product.ID).Updates(map[string]interface{}{
		"name":       product.Name,
		"price":      product.Price,
		"updated_at": product.UpdatedAt,
	}).Error
}

func (r *productRepository) Delete(id uint) error {
	return config.DB.Delete(&models.Product{}, "id = ?", id).Error
}