package seeders

import (
	"github.com/hoyirul/go-starter-kit/internal/models"

	"gorm.io/gorm"
)

type ProductSeeder struct{}

func (s *ProductSeeder) Seed(db *gorm.DB) error {
	products := []models.Product{
		{UserID: "5bb40e88-627e-461d-8f3a-9e8453c64f5a", Name: "Product A", Price: 10.0},
		{UserID: "5bb40e88-627e-461d-8f3a-9e8453c64f5a", Name: "Product B", Price: 20.0},
		{UserID: "38083655-a6bf-4db3-adb5-47326a70e3ce", Name: "Product C", Price: 15.5},
	}

	for _, p := range products {
		if err := db.Create(&p).Error; err != nil {
			return err
		}
	}

	return nil
}

func (s *ProductSeeder) Unseed(db *gorm.DB) error {
	return db.Exec("DELETE FROM products").Error
}
