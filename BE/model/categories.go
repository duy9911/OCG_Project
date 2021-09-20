package model

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryName string    `json:"category_name"`
	CollectionID uint      `json:"collection_id"`
	Products     []Product `gorm:"many2many:category_products;"`
}

type CategoryProduct struct {
	gorm.Model
	ProductID  uint `json:"product_id"`
	CategoryID uint `json:"category_id"`
}
