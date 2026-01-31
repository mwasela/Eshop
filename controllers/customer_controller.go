package controllers

import (
	"Eshop/models"
	"Eshop/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCustomer(c *gin.Context) {
	var input struct {
		Name    string `json:"name" binding:"required"`
		Email   string `json:"email" binding:"required,email"`
		Phone   string `json:"phone" binding:"required"`
		Address string `json:"address"`
		Type    int    `json:"type" binding:"required"` // 1 = Walk in customer, 2 = Contractor
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer := models.Customers{
		Name:    input.Name,
		Email:   input.Email,
		Phone:   input.Phone,
		Address: input.Address,
		Type:    input.Type,
	}

	if err := config.DB.Create(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer or email/phone exists"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Customer created successfully", "customer": customer})
}


func GetCustomers(c *gin.Context) {
	var customers []models.Customers
	if err := config.DB.Find(&customers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve customers"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"customers": customers})
}


func GetCustomerByID(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customers

	if err := config.DB.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"customer": customer})
}

func UpdateCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customers

	if err := config.DB.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	var input struct {
		Name    string `json:"name"`
		Email   string `json:"email" binding:"email"`
		Phone   string `json:"phone"`
		Address string `json:"address"`
		Type    int    `json:"type"` // 1 = Walk in customer, 2 = Contractor
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedCustomer := models.Customers{
		Name:    input.Name,
		Email:   input.Email,
		Phone:   input.Phone,
		Address: input.Address,
		Type:    input.Type,
	}

	if err := config.DB.Model(&customer).Updates(updatedCustomer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update customer or email/phone exists"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer updated successfully", "customer": customer})
}

func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customers

	if err := config.DB.First(&customer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}
	
	if err := config.DB.Delete(&customer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete customer"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully"})
}
