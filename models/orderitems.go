package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	//generate oder id during order creation
	OrderID     uint     `json:"order_id" gorm:"column:eshop_order_item_order_id;not null"`
	ProductID   uint     `json:"product_id" gorm:"column:eshop_order_item_product_id;not null"`
	Product     Product `json:"product" gorm:"foreignKey:ProductID"`
	Quantity    int      `json:"quantity" gorm:"column:eshop_order_item_quantity;not null"`
	UnitPrice   float64  `json:"unit_price" gorm:"column:eshop_order_item_unit_price;not null"`
	TotalPrice  float64  `json:"total_price" gorm:"column:eshop_order_item_total_price;not null"`
	
}
