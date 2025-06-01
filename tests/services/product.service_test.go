package tests

import (
	"errors"
	"testing"

	"github.com/hoyirul/go-starter-kit/internal/models"
	"github.com/hoyirul/go-starter-kit/internal/services"
	mocks "github.com/hoyirul/go-starter-kit/mocks/services"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// TestGetProducts tests the GetProducts method of ProductService.
func TestGetProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	productService := services.NewProductService(mockRepo)

	expectedProducts := []models.ProductResponse{
		{ID: 1, Name: "Product 1", Price: 100.0},
		{ID: 2, Name: "Product 2", Price: 200.0},
	}

	t.Run("Products found", func(t *testing.T) {
		mockRepo.EXPECT().
			FindAll(gomock.Any(), "").
			Return(expectedProducts, nil, nil)

		products, pagination, err := productService.GetProducts(nil, "")

		assert.NoError(t, err)
		assert.Equal(t, expectedProducts, products)
		assert.Nil(t, pagination)
	})

	t.Run("Error fetching products", func(t *testing.T) {
		mockRepo.EXPECT().
			FindAll(gomock.Any(), "").
			Return(nil, nil, errors.New("fetch error"))

		products, pagination, err := productService.GetProducts(nil, "")

		assert.Error(t, err)
		assert.Nil(t, products)
		assert.Nil(t, pagination)
	})
}

// TestGetProductByID tests the GetProduct method of ProductService.
func TestGetProductByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	productService := services.NewProductService(mockRepo)

	expectedProduct := &models.ProductResponse{
		ID:    1,
		Name:  "Test Product",
		Price: 100.0,
	}

	t.Run("Product found", func(t *testing.T) {
		mockRepo.EXPECT().
			FindByID(uint(1)).
			Return(expectedProduct, nil)

		product, err := productService.GetProduct(1)

		assert.NoError(t, err)
		assert.Equal(t, expectedProduct, product)
	})

	t.Run("Product not found", func(t *testing.T) {
		mockRepo.EXPECT().
			FindByID(uint(2)).
			Return(nil, errors.New("not found"))

		product, err := productService.GetProduct(2)

		assert.Error(t, err)
		assert.Nil(t, product)
	})
}

// TestCreateProduct tests the CreateProduct method of ProductService.
func TestCreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	productService := services.NewProductService(mockRepo)

	newProduct := &models.Product{
		Name:  "New Product",
		Price: 50.0,
	}

	t.Run("Create product successfully", func(t *testing.T) {
		mockRepo.EXPECT().
			Create(newProduct).
			Return(nil)

		err := productService.CreateProduct(newProduct)

		assert.NoError(t, err)
	})

	t.Run("Create product with error", func(t *testing.T) {
		mockRepo.EXPECT().
			Create(newProduct).
			Return(errors.New("creation error"))

		err := productService.CreateProduct(newProduct)

		assert.Error(t, err)
	})
}

// TestUpdateProduct tests the UpdateProduct method of ProductService.
func TestUpdateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	productService := services.NewProductService(mockRepo)

	updatedProduct := &models.Product{
		ID:    1,
		Name:  "Updated Product",
		Price: 75.0,
	}

	t.Run("Update product successfully", func(t *testing.T) {
		mockRepo.EXPECT().
			Update(updatedProduct).
			Return(nil)

		err := productService.UpdateProduct(updatedProduct)

		assert.NoError(t, err)
	})

	t.Run("Update product with invalid ID", func(t *testing.T) {
		invalidProduct := &models.Product{
			Name:  "Invalid Product",
			Price: 30.0,
		}

		err := productService.UpdateProduct(invalidProduct)

		assert.Error(t, err)
	})
}

// TestDeleteProduct tests the DeleteProduct method of ProductService.
func TestDeleteProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepository(ctrl)
	productService := services.NewProductService(mockRepo)

	t.Run("Delete product successfully", func(t *testing.T) {
		mockRepo.EXPECT().
			Delete(uint(1)).
			Return(nil)

		err := productService.DeleteProduct(1)

		assert.NoError(t, err)
	})

	t.Run("Delete product with invalid ID", func(t *testing.T) {
		err := productService.DeleteProduct(0)

		assert.Error(t, err)
	})
}
