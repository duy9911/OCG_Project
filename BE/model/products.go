package model

import (
	"backend/dbconfig"
	"fmt"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name             string            `json: "name"`
	Description      string            `json: "description`
	ProductVariances []ProductVariance `gorm:"foreignKey:ProductID"`
	Images           []Image           `gorm:"foreignKey:ProductID"`
	CategoryProducts []CategoryProduct `gorm:"foreignKey:ProductID"`
}
type ProductVariance struct {
	gorm.Model
	ProductID   uint   `json: "product_id"`
	Color       string `json: "color"`
	Size        string `json: "size"`
	OrderDetail OrderDetail
}

type CreateProduct struct {
	ID          uint   `json: "product_id"`
	Name        string `json: "name"`
	Description string `json: "description"`
	Color       string `json: "color"`
	Size        string `json: "size"`
	Image       string `json: "image"`
}

var product []CreateProduct

func AllProduct() ([]CreateProduct, error) {
	rows, err := dbconfig.Database.Debug().Table("products").Select(" products.id, products.name, products.description, product_variances.color, product_variances.size, images.image").Joins("inner join  product_variances on products.id=product_variances.product_id inner join images on products.id=images.product_id").Rows()
	fmt.Println(rows)
	defer rows.Close()
	for rows.Next() {
		dbconfig.Database.ScanRows(rows, &product)
	}
	return product, err
}
