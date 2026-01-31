package models

import "gorm.io/gorm"

type Customers struct {
    gorm.Model
    Name    string `json:"name" gorm:"column:eshop_customer_name;not null"`
    Email   string `json:"email" gorm:"column:eshop_customer_email;unique;not null"`
    Phone   string `json:"phone" gorm:"column:eshop_customer_phone;unique;not null"`
    Address string `json:"address" gorm:"column:eshop_customer_address"`
    Type    int    `json:"type" gorm:"column:eshop_customer_type"` // 1 = Walk in customer, 2 = Contractor
}