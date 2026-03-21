package models

import "gorm.io/gorm"

type Storetype struct {
	gorm.Model
	Name        string `json:"storetype_name" gorm:"column:eshop_storetype_name;unique;not null"`
	Description string `json:"storetype_description" gorm:"column:eshop_storetype_description"`
	Status      int    `json:"storetype_status" gorm:"column:eshop_storetype_status;default:1"`
}

