package models

import "gorm.io/gorm"


type SalesOrder struct {
	gorm.Model
	CustomerID   uint      `json:"customer_id" gorm:"not null"`
	Customer     Customers `json:"customer" gorm:"foreignKey:CustomerID"`
	OrderDate    string    `json:"order_date" gorm:"column:eshop_salesorder_orderdate;not null"`
	TotalAmount  float64   `json:"total_amount" gorm:"column:eshop_salesorder_totalamount;not null"`
	TaxAmount    float64   `json:"tax_amount" gorm:"column:eshop_salesorder_taxamount;not null"`
	Status       string    `json:"status" gorm:"column:eshop_salesorder_status;not null"` // e.g., Pending, Completed, Cancelled
}