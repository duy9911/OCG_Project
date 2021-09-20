package model

import (
	"backend/config"
	"fmt"

	"gorm.io/gorm"
)

type Page struct {
	gorm.Model
	PageName    string       `json:"page_name"`
	Collections []Collection `gorm:"foreignKey:PageID"`
}

func OnePageCollections(id string) (Page, error) {
	rt := Page{}
	if err := config.Database.Where("id = ? ", id).First(&rt).Error; err != nil {
		fmt.Println(err.Error())
	}
	config.Database.Model(&rt).Association("Collections").Find(&rt.Collections)
	return rt, nil
}

var pages []Page

//Preload collection first and then match with pages
func AllPageCollections() ([]Page, error) {
	err := config.Database.Preload("Collections").Find(&pages).Error // SELECT * FROM pages;
	return pages, err                                                // SELECT * FROM colelctions WHERE colelctions.pages_id IN (1,2,3,4);

}
