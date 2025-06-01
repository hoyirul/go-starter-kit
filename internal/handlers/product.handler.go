package handlers

import (
	"net/http"

	"github.com/hoyirul/go-starter-kit/internal/models"
	"github.com/hoyirul/go-starter-kit/internal/services"
	"github.com/hoyirul/go-starter-kit/utils"

	validation "github.com/go-ozzo/ozzo-validation"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Service services.ProductService
}

func NewProductHandler(service services.ProductService) *ProductHandler {
	return &ProductHandler{Service: service}
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	search := c.Query("search")
	products, pagination, err := h.Service.GetProducts(c, search)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to fetch products")
		return
	}

	if len(products) == 0 {
		utils.RespondWithSuccess(c, http.StatusOK, "No products found", gin.H{
			"data":       []models.Product{},
			"pagination": pagination,
		})
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Products fetched successfully", gin.H{
		"data":       products,
		"pagination": pagination,
	})
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	productID, err := utils.ParseUint(id)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product, err := h.Service.GetProduct(productID)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Product not found")
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Product fetched successfully", product)
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var request models.ProductRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required, validation.Length(1, 100)),
		validation.Field(&request.Price, validation.Required, validation.Min(0.01)),
	); err != nil {
		errors := make(map[string]string)
		for field, err := range err.(validation.Errors) {
			errors[field] = err.Error()
		}
		utils.RespondWithValidationErrors(c, http.StatusBadRequest, errors)
		return
	}

	user, exists := c.Get("user")
	if !exists {
		utils.RespondWithError(c, http.StatusUnauthorized, "User not authenticated")
		return
	}

	u, ok := user.(*models.User)
	if !ok {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to get user information")
		return
	}
	
	product := models.Product{
		UserID: u.ID,
		Name:   request.Name,
		Price:  request.Price,
		CreatedAt: utils.GetCurrentTime(),
		UpdatedAt: utils.GetCurrentTime(),
	}

	if err := h.Service.CreateProduct(&product); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create product")
		return
	}

	utils.RespondWithSuccess(c, http.StatusCreated, "Product created successfully", product)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	productID, err := utils.ParseUint(id)
	var request models.ProductRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required, validation.Length(1, 100)),
		validation.Field(&request.Price, validation.Required, validation.Min(0.01)),
	); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload: "+err.Error())
		return
	}	

	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid product ID")
		return
	}
	
	existingProduct, err := h.Service.GetProduct(productID)
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Product not found")
		return
	}

	if existingProduct == nil {
		utils.RespondWithError(c, http.StatusNotFound, "Product not found")
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, exists := c.Get("user")
	if !exists {
		utils.RespondWithError(c, http.StatusUnauthorized, "User not authenticated")
		return
	}

	u, ok := user.(*models.User)
	if !ok {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to get user information")
		return
	}

	product := models.Product{
		ID:     productID,
		UserID: u.ID,
		Name:   request.Name,
		Price:  request.Price,
		CreatedAt: existingProduct.CreatedAt,
		UpdatedAt: utils.GetCurrentTime(),
	}

	if err := h.Service.UpdateProduct(&product); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to update product")
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Product updated successfully", product)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	productID, err := utils.ParseUint(id)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	if err := h.Service.DeleteProduct(productID); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to delete product")
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Product deleted successfully", nil)
}