package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string `json:"name" gorm:"column:eshop_category_name;unique;not null"`
	Description string `json:"description" gorm:"column:eshop_category_description"`
}