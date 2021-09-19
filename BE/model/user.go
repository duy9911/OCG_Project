package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}
