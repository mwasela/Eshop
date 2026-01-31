package config

import (
	"Eshop/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "michael:12345@tcp(127.0.0.1:3306)/go_eshop?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	err = database.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Customers{},
		&models.Inventory{},
		&models.OrderItem{},
		&models.Pricelist{},
		&models.Product{},
		&models.SalesOrder{},
		&models.Supplier{},
	)

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database
}
