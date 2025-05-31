package models

import "time"

type (
	Product struct {
		ID          uint		  `gorm:"primaryKey;autoIncrement" json:"id"`
		UserID      string    `json:"user_id"`
		Name        string    `json:"name"`
		Price			  float64   `json:"price"`
		CreatedAt	  time.Time `gorm:"autoCreateTime" json:"created_at"`
		UpdatedAt		time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	}

	ProductRequest struct {
    Name  string  `json:"name" validate:"required,min=1,max=100"`
    Price float64 `json:"price" validate:"required,gt=0"`
	}

	ProductResponse struct {
		ID        uint    	`json:"id"`
		UserID    string  	`json:"user_id"`
		Name      string  	`json:"name"`
		Price     float64 	`json:"price"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		User 			User 			`gorm:"foreignKey:UserID" json:"user"`
	}
)

func (p *ProductResponse) TableName() string {
	return "products"
}