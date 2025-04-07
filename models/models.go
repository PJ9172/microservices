package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:email gorm:"unique"`
	Password string `json:password`
}

type Order struct {
	gorm.Model
	Order_id string `json:order_id`
	Email    string `json:email`
	Payment  string `json:payment`
}
