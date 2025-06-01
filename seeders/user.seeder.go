package seeders

import (
	"github.com/hoyirul/go-starter-kit/internal/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserSeeder struct{}

func (s *UserSeeder) Seed(db *gorm.DB) error {
	users := []models.User{
		{ID: "5bb40e88-627e-461d-8f3a-9e8453c64f5a", Name: "John Doe", Email: "john@mail.com"},
		{ID: "38083655-a6bf-4db3-adb5-47326a70e3ce", Name: "Alice Johnson", Email: "jane@mail.com"},
	}

	for i := range users {
		hashed, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		users[i].Password = string(hashed)

		if err := db.FirstOrCreate(&users[i], models.User{ID: users[i].ID}).Error; err != nil {
			return err
		}
	}

	return nil
}

func (s *UserSeeder) Unseed(db *gorm.DB) error {
	return db.Exec("DELETE FROM users").Error
}
