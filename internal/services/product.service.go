package services

import (
	"go-starter-kit/internal/models"
	"go-starter-kit/internal/repository"
	"go-starter-kit/utils"

	"github.com/gin-gonic/gin"
)

type ProductService interface {
	GetProducts(c *gin.Context, search string) ([]models.ProductResponse, *utils.Pagination, error)
	GetProduct(id uint) (*models.ProductResponse, error)
	CreateProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
	DeleteProduct(id uint) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo}
}

func (s *productService) GetProducts(c *gin.Context, search string) ([]models.ProductResponse, *utils.Pagination, error) {
	products, pagination, err := s.repo.FindAll(c, search)
	if err != nil {
		return nil, nil, err
	}
	return products, pagination, nil
}

func (s *productService) GetProduct(id uint) (*models.ProductResponse, error) {
	product, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *productService) CreateProduct(product *models.Product) error {
	return s.repo.Create(product)
}

func (s *productService) UpdateProduct(product *models.Product) error {
	if product.ID == 0 {
		return utils.ErrInvalidID // Ensure ID is valid for update
	}
	return s.repo.Update(product)
}

func (s *productService) DeleteProduct(id uint) error {
	if id == 0 {
		return utils.ErrInvalidID // Ensure ID is valid for deletion
	}
	return s.repo.Delete(id)
}