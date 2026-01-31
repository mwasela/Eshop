package models

import "gorm.io/gorm"

type Inventory struct {
	gorm.Model
	ProductID uint   `json:"product_id" gorm:"column:eshop_inventory_product_id;not null"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
	BinLocation         string   `json:"bin_location" gorm:"column:eshop_inventory_bin_location;unique;not null"`
	StockQuantity     int      `json:"stock_quantity" gorm:"column:eshop_inventory_stock_quantity;not null"`
	Threshold        float64  `json:"threshold" gorm:"column:eshop_inventory_threshold;not null"`
}