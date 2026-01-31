package models

import "gorm.io/gorm"

type Pricelist struct {
	gorm.Model
	ProductID uint    `json:"product_id" gorm:"not null"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	Costprice float64 `json:"costprice" gorm:"column:eshop_pricelist_costprice;not null"`
	Wholsaleprice float64 `json:"wholsaleprice" gorm:"column:eshop_pricelist_wholsaleprice;not null"`
	Retailprice   float64 `json:"retailprice" gorm:"column:eshop_pricelist_retailprice;not null"`
}