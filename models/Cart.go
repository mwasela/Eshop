package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	CartID     string  `json:"cart_id" gorm:"column:eshop_cart_id;"`
	ProductID  uint    `json:"product_id" gorm:"column:eshop_product_id;not null"`
	Product    Product `json:"product" gorm:"foreignKey:ProductID"`
	Quantity   int     `json:"quantity" gorm:"column:eshop_quantity;not null"`
	UnitPrice  float64 `json:"unit_price" gorm:"column:eshop_unit_price;not null"`
	TotalPrice float64 `json:"total_price" gorm:"column:eshop_total_price;not null"`
	UserID    *uint   `json:"user_id" gorm:"column:eshop_user_id"`
	User      *User   `json:"user" gorm:"foreignKey:UserID"`
}