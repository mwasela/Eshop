package models

import "gorm.io/gorm"

type Companies struct {
	gorm.Model
	CompanyName string `json:"company_name" gorm:"column:eshop_company_name;unique;not null"`
	Address     string `json:"address" gorm:"column:eshop_company_address"`
	Phone       string `json:"phone" gorm:"column:eshop_company_phone"`
	Email       string `json:"email" gorm:"column:eshop_company_email"`
	Status      int    `json:"status" gorm:"column:eshop_company_status;default:1"`
}


