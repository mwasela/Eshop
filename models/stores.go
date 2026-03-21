package models

import "gorm.io/gorm"

type Stores struct {
	gorm.Model
	StoreName string `json:"store_name" gorm:"column:eshop_store_name;unique;not null"`
	LocationID uint   `json:"location_id" gorm:"column:eshop_store_location_id;not null"`
	Location   Location `json:"location" gorm:"foreignKey:LocationID"`
	Area 	string `json:"area" gorm:"column:eshop_store_area"`
	CompanyID uint   `json:"company_id" gorm:"column:eshop_store_company_id;not null"`
	Companies    Companies `json:"company" gorm:"foreignKey:CompanyID"`
	StoretypeID   uint `json:"storetype_id" gorm:"column:eshop_store_type_id;not null"`
	Storetype   Storetype `json:"storetype" gorm:"foreignKey:StoretypeID"`
	Status	int    `json:"status" gorm:"column:eshop_store_status;default:1"`

}
