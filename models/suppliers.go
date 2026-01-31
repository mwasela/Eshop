package models

import "gorm.io/gorm"

type Supplier struct {
	gorm.Model
    Contact     string  `json:"contact" gorm:"eshop_supplier_contact;not null"`
	Location	 string  `json:"location" gorm:"eshop_supplier_location"`
	SupplierName string  `json:"supplier_name" gorm:"eshop_supplier_name"`
}