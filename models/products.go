package models

import "gorm.io/gorm"

type Product struct {
    gorm.Model
    Name        string   `json:"name" gorm:"column:eshop_product_name;not null"`
    SKU         string   `json:"sku" gorm:"column:eshop_product_sku;unique;not null"`	
    OEM         string   `json:"oem" gorm:"column:eshop_product_oem"`
    Description string   `json:"description" gorm:"column:eshop_product_description"`
    CategoryID  uint     `json:"category_id" gorm:"column:eshop_product_category_id"`
    Category    Category `json:"category" gorm:"foreignKey:CategoryID"`
}