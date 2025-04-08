package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:email gorm:"unique"`
	Password string `json:password`
}

type Purchase struct {
	gorm.Model
	Email    string   `json:email`
	Status   string   `json:status`
	Products []string `gorm:"-"`
	RawData  string   `json:rawdata`
}
