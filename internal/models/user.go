package models

import "time"

type (
	User struct {
		ID				string `gorm:"primaryKey" json:"id"`
		Name 			string `json:"name"`
		Email     string `gorm:"unique" json:"email"`
		Password  string `json:"password"`
		CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
		UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	}

	UserRequest struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	UserResponse struct {
		ID         string    `json:"id"`
		Name       string    `json:"name"`
		Email      string    `json:"email"`
		CreatedAt  time.Time `json:"created_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	}
)

func (u *UserResponse) TableName() string {
	return "users"
}