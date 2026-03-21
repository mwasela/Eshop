package models

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	Name    string `json:"location_name" gorm:"column:eshop_location_name;unique;not null"`
	Country string `json:"location_country" gorm:"column:eshop_location_country"`
}