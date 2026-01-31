package controllers

import(
	"Eshop/models"
	"Eshop/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrderItem(c *gin.Context) {
	var input struct {
		//generate orderid
		OrderID     uint  `json:"order_id" binding:"required"`
		ProductID   uint  `json:"product_id" binding:"required"`
		Quantity    int     `json:"quantity" binding:"required"`
		UnitPrice   float64 `json:"unit_price" binding:"required"`
		TotalPrice  float64 `json:"total_price" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	orderItem := models.OrderItem{
		OrderID:    input.OrderID,
		ProductID:  input.ProductID,
		Quantity:   input.Quantity,
		UnitPrice:  input.UnitPrice,
		TotalPrice: input.TotalPrice,
	}

	if err := config.DB.Create(&orderItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order item"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Order item created successfully", "order_item": orderItem})
}

func GetOrderItems(c *gin.Context) {
	var orderItems []models.OrderItem
	if err := config.DB.Find(&orderItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve order items"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order_items": orderItems})
}


func GetOrderItemByID(c *gin.Context) {
	id := c.Param("id")
	var orderItem models.OrderItem

	if err := config.DB.First(&orderItem, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order item not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order_item": orderItem})
}

func UpdateOrderItem(c *gin.Context) {
	id := c.Param("id")
	var orderItem models.OrderItem

	if err := config.DB.First(&orderItem, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order item not found"})
		return
	}

	var input struct {
		OrderID     uint  `json:"order_id"`
		ProductID   uint  `json:"product_id"`
		Quantity    int     `json:"quantity"`
		UnitPrice   float64 `json:"unit_price"`
		TotalPrice  float64 `json:"total_price"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedOrderItem := models.OrderItem{
		OrderID:    input.OrderID,
		ProductID:  input.ProductID,
		Quantity:   input.Quantity,
		UnitPrice:  input.UnitPrice,
		TotalPrice: input.TotalPrice,
	}

	if err := config.DB.Model(&orderItem).Updates(updatedOrderItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order item updated successfully", "order_item": orderItem})
}


func DeleteOrderItem(c *gin.Context) {
	id := c.Param("id")
	var orderItem models.OrderItem

	if err := config.DB.First(&orderItem, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order item not found"})
		return
	}

	if err := config.DB.Delete(&orderItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order item deleted successfully"})
}

