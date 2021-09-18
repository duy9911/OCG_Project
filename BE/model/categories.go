package model

import (
	"backend/dbconfig"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryName     string `json:"category_name"`
	CollectionID     uint   `json:"collection_id"`
	CategoryProducts []CategoryProduct
}

type CategoryProduct struct {
	gorm.Model
	ProductID  uint `json: "product_id `
	CategoryID uint `json: "category_id"`
}

type Collection struct {
	gorm.Model
	Image          string     `json:"image"`
	CollectionName string     `json:"collection_name"`
	PageID         uint       `json:"page_id"`
	Categories     []Category `gorm:"foreignKey:CollectionID"`
}

type Page struct {
	gorm.Model
	PageName   string       `json:"page_name"`
	Collection []Collection `gorm:"foreignKey:PageID"`
}

type CreatePage struct {
	ID             uint   `json:"id"`
	PageName       string `json:"page_name"`
	CollectionName string `json:"collection_name"`
}

var createPages []CreatePage

//tao mot struct de output api home page/collection
// type output struct{
// ID uint
// PageName
// Collections  []Collection
// }
// lap qua trung rows dien 2 variable id va name sau do tao 1 slice va dien collection_id va collection_name
// trong moi moi rows loop qua page_id is exist,
func AllPage() ([]CreatePage, error) {
	rows, err := dbconfig.Database.Debug().Table("pages").Select("pages.id, pages.page_name, collections.collection_name").Joins("inner join collections on pages.id = collections.page_id").Rows()
	defer rows.Close()
	for rows.Next() {
		dbconfig.Database.ScanRows(rows, &createPages)
	}
	return createPages, err
}

func OnePage(id string) (CreatePage, error) {
	createPage := CreatePage{}
	rows, err := dbconfig.Database.Table("pages").Select("pages.id, pages.page_name, collections.collection_name").Joins("inner join collections on pages.id=collections.page_id").Where("pages.id=?", id).Find(&createPage).Rows()
	defer rows.Close()
	for rows.Next() {
		dbconfig.Database.ScanRows(rows, &createPage)
	}
	return createPage, err
}
