package controllers

import (
	"Eshop/models"
	"Eshop/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateSupplier(c *gin.Context) {
	var input struct {
		Contact      string `json:"contact" binding:"required"`
		Location     string `json:"location"`
		SupplierName string `json:"supplier_name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	supplier := models.Supplier{
		Contact:      input.Contact,
		Location:     input.Location,
		SupplierName: input.SupplierName,
	}

	if err := config.DB.Create(&supplier).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create supplier"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Supplier created successfully", "supplier": supplier})
}

func GetSuppliers(c *gin.Context) {
	var suppliers []models.Supplier
	if err := config.DB.Find(&suppliers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve suppliers"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"suppliers": suppliers})
}

func GetSupplierByID(c *gin.Context) {
	id := c.Param("id")
	var supplier models.Supplier

	if err := config.DB.First(&supplier, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Supplier not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"supplier": supplier})
}

func UpdateSupplier(c *gin.Context) {
	id := c.Param("id")
	var supplier models.Supplier

	if err := config.DB.First(&supplier, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Supplier not found"})
		return
	}

	var input struct {
		Contact      string `json:"contact"`
		Location     string `json:"location"`
		SupplierName string `json:"supplier_name"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	supplier.Contact = input.Contact
	supplier.Location = input.Location
	supplier.SupplierName = input.SupplierName

	if err := config.DB.Save(&supplier).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update supplier"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Supplier updated successfully", "supplier": supplier})
}	


func DeleteSupplier(c *gin.Context) {
	id := c.Param("id")
	var supplier models.Supplier

	if err := config.DB.First(&supplier, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Supplier not found"})
		return
	}

	if err := config.DB.Delete(&supplier).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete supplier"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Supplier deleted successfully"})
}	