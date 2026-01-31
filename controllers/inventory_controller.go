package controllers

import (
	"Eshop/models"
	"Eshop/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateInventoryItem(c *gin.Context) {
	var input struct {
		ProductID      uint  `json:"product_id" binding:"required"`
		StockQuantity  int     `json:"stock_quantity" binding:"required"`
		Threshold   float64  `json:"threshold" binding:"required"`
		BinLocation    string  `json:"bin_location" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	inventoryItem := models.Inventory{
		ProductID:     input.ProductID,
		StockQuantity: input.StockQuantity,
		Threshold: input.Threshold,
		BinLocation:   input.BinLocation,
	}
	
	if err := config.DB.Create(&inventoryItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create inventory item or bin location exists"})
		return
	}
	
	// Reload inventory item with product and category
	config.DB.Preload("Product.Category").Preload("Product").First(&inventoryItem, inventoryItem.ID)
	
	c.JSON(http.StatusOK, gin.H{"message": "Inventory item created successfully", "inventory_item": inventoryItem})
}

func GetInventoryItems(c *gin.Context) {
	var inventoryItems []models.Inventory
	if err := config.DB.Preload("Product.Category").Preload("Product").Find(&inventoryItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve inventory items"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"inventory_items": inventoryItems})
}


func GetInventoryItemByID(c *gin.Context) {
	id := c.Param("id")
	var inventoryItem models.Inventory

	if err := config.DB.Preload("Product.Category").Preload("Product").First(&inventoryItem, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory item not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"inventory_item": inventoryItem})
}

func UpdateInventoryItem(c *gin.Context) {
	id := c.Param("id")
	var inventoryItem models.Inventory

	if err := config.DB.First(&inventoryItem, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory item not found"})
		return
	}
	
	var input struct {
		ProductID      uint  `json:"product_id"`
		StockQuantity  int     `json:"stock_quantity"`
		Threshold   float64  `json:"threshold"`
		BinLocation    string  `json:"bin_location"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.ProductID != 0 {
		inventoryItem.ProductID = input.ProductID
	}
	if input.StockQuantity != 0 {
		inventoryItem.StockQuantity = input.StockQuantity
	}
	if input.Threshold != 0 {
		inventoryItem.Threshold = input.Threshold
	}
	if input.BinLocation != "" {
		inventoryItem.BinLocation = input.BinLocation
	}

	if err := config.DB.Save(&inventoryItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update inventory item"})
		return
	}

	// Reload inventory item with product and category
	config.DB.Preload("Product.Category").Preload("Product").First(&inventoryItem, id)

	c.JSON(http.StatusOK, gin.H{"message": "Inventory item updated successfully", "inventory_item": inventoryItem})
}

func DeleteInventoryItem(c *gin.Context) {
	id := c.Param("id")
	var inventoryItem models.Inventory

	if err := config.DB.First(&inventoryItem, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory item not found"})
		return
	}

	if err := config.DB.Delete(&inventoryItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete inventory item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inventory item deleted successfully"})
}	


