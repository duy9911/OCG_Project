package model

import "gorm.io/gorm"

type Image struct {
	gorm.Model
	ProductID uint   `json:"product_id"`
	Image     string `json:"image"`
}
