package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func init() {
	var err error
	dsn := "root:duyngo99@tcp(127.0.0.1:3306)/decorshop2?charset=utf8mb4&parseTime=True&loc=Local"
	Database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successful.")
}
