package Models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Products struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Price       int    `gorm:"not null"`
	Description string `gorm:"not null"`
}

type ProductsResponse struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Price       int       `json:"price"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
