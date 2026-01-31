package controllers

import (
	"Eshop/models"
	"Eshop/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePriceListItem(c *gin.Context) {
	var input struct {
		ProductID   uint     `json:"product_id" binding:"required"`
		Costprice   float64 `json:"costprice" binding:"required"`
		Wholsaleprice float64 `json:"wholsaleprice" binding:"required"`
		Retailprice  float64 `json:"retailprice" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	priceListItem := models.Pricelist{
		ProductID:    input.ProductID,
		Costprice:    input.Costprice,
		Wholsaleprice: input.Wholsaleprice,
		Retailprice:  input.Retailprice,
	}

	if err := config.DB.Create(&priceListItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create price list item"})
		return
	}
	
	// Reload price list item with product and category
	config.DB.Preload("Product.Category").Preload("Product").First(&priceListItem, priceListItem.ID)
	
	c.JSON(http.StatusOK, gin.H{"message": "Price list item created successfully", "price_list_item": priceListItem})
}

func GetPriceListItems(c *gin.Context) {
	var priceListItems []models.Pricelist
	if err := config.DB.Preload("Product.Category").Preload("Product").Find(&priceListItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve price list items"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"price_list_items": priceListItems})
}


func GetPriceListItemByID(c *gin.Context) {
	id := c.Param("id")
	var priceListItem models.Pricelist

	if err := config.DB.Preload("Product.Category").Preload("Product").First(&priceListItem, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Price list item not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"price_list_item": priceListItem})
}

func UpdatePriceListItem(c *gin.Context) {
	id := c.Param("id")
	var priceListItem models.Pricelist

	if err := config.DB.First(&priceListItem, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Price list item not found"})
		return
	}

	var input struct {
		Costprice    float64 `json:"costprice"`
		Wholsaleprice float64 `json:"wholsaleprice"`
		Retailprice  float64 `json:"retailprice"`
	}
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	priceListItem.Costprice = input.Costprice
	priceListItem.Wholsaleprice = input.Wholsaleprice
	priceListItem.Retailprice = input.Retailprice

	if err := config.DB.Save(&priceListItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update price list item"})
		return
	}

	// Reload price list item with product and category
	config.DB.Preload("Product.Category").Preload("Product").First(&priceListItem, id)

	c.JSON(http.StatusOK, gin.H{"message": "Price list item updated successfully", "price_list_item": priceListItem})
}

func DeletePriceListItem(c *gin.Context) {
	id := c.Param("id")
	var priceListItem models.Pricelist

	if err := config.DB.First(&priceListItem, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Price list item not found"})
		return
	}

	if err := config.DB.Delete(&priceListItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete price list item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Price list item deleted successfully"})
}
