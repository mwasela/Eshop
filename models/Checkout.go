package models

import "gorm.io/gorm"

type Checkout struct {
	gorm.Model
	CartID                string  `json:"cart_id" gorm:"column:eshop_cart_id;size:41;not null"`
	PaymentType           int     `json:"payment_type" gorm:"column:eshop_payment_type;not null;check:eshop_payment_type IN (1,2)"`
	ClientFname           string  `json:"client_fname" gorm:"column:eshop_client_fname;size:20;not null"`
	ClientLname           string  `json:"client_lname" gorm:"column:eshop_client_lname;size:20;not null"`
	DeliveryCity          string  `json:"delivery_city" gorm:"column:eshop_delivery_city;size:20;not null"`
	DeliveryStreetAddress string  `json:"delivery_street_address" gorm:"column:eshop_delivery_street_address;size:60;not null"`
	DeliveryCountry       string  `json:"delivery_country" gorm:"column:eshop_delivery_country;size:20;not null"`
	DeliveryNotes         string  `json:"delivery_notes" gorm:"column:eshop_delivery_notes;size:100"`
	ClientPhone           string  `json:"client_phone" gorm:"column:eshop_client_phone;size:14"`
	TotalAmount           float64 `json:"total_amount" gorm:"column:eshop_total_amount;not null"`
	PaymentFlag           int     `json:"payment_flag" gorm:"column:eshop_payment_flag;not null;default:1;check:eshop_payment_flag IN (0,1)"`
	CompanyID			  uint    `json:"company_id" gorm:"column:eshop_company_id;not null"`
	Company				  Company `json:"company" gorm:"foreignKey:CompanyID"`
}