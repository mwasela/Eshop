package controllers

import (
	"Eshop/models"
	"Eshop/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateSalesOrder(c *gin.Context) {
	var input struct{
		CustomerID uint    `json:"customer_id" binding:"required"`
		OrderDate   string  `json:"order_date" binding:"required"`
		TotalAmount float64 `json:"total_amount" binding:"required"`
		TaxAmount   float64 `json:"tax_amount" binding:"required"`
		Status      string  `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	salesOrder := models.SalesOrder{
		CustomerID:  input.CustomerID,
		OrderDate:   input.OrderDate,
		TotalAmount: input.TotalAmount,
		TaxAmount:   input.TaxAmount,
		Status:      input.Status,
	}
	
	if err := config.DB.Create(&salesOrder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create sales order"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Sales order created successfully", "sales_order": salesOrder})
}

func GetSalesOrders(c *gin.Context) {
	var salesOrders []models.SalesOrder
	if err := config.DB.Find(&salesOrders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve sales orders"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"sales_orders": salesOrders})
}

func GetSalesOrderByID(c *gin.Context) {
	id := c.Param("id")
	var salesOrder models.SalesOrder

	if err := config.DB.First(&salesOrder, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sales order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"sales_order": salesOrder})
}

func UpdateSalesOrder(c *gin.Context) {
	id := c.Param("id")
	var salesOrder models.SalesOrder

	if err := config.DB.First(&salesOrder, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sales order not found"})
		return
	}

	var input struct{
		CustomerID uint    `json:"customer_id"`
		OrderDate   string  `json:"order_date"`
		TotalAmount float64 `json:"total_amount"`
		TaxAmount   float64 `json:"tax_amount"`
		Status      string  `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	salesOrder.CustomerID = input.CustomerID
	salesOrder.OrderDate = input.OrderDate
	salesOrder.TotalAmount = input.TotalAmount
	salesOrder.TaxAmount = input.TaxAmount
	salesOrder.Status = input.Status

	if err := config.DB.Save(&salesOrder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update sales order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sales order updated successfully", "sales_order": salesOrder})
}

func DeleteSalesOrder(c *gin.Context) {
	id := c.Param("id")
	var salesOrder models.SalesOrder

	if err := config.DB.First(&salesOrder, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sales order not found"})
		return
	}

	if err := config.DB.Delete(&salesOrder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete sales order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Sales order with ID %s deleted successfully", id)})
}
