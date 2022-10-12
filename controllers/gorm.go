package controllers

import "github.com/jinzhu/gorm"

type Person struct {
	gorm.Model
	First_name string
	Last_name  string
}

type InDB struct {
	DB *gorm.DB
}
