package controllers

import (
	"Eshop/models"
	"Eshop/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateStoretype(c *gin.Context) {
	var input struct {
		Name        string `json:"storetype_name" binding:"required"`
		Description string `json:"storetype_description"`
		Status      int    `json:"storetype_status"`
	}
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storetype := models.Storetype{
		Name:        input.Name,
		Description: input.Description,
		Status:      input.Status,
	}

	if err := config.DB.Create(&storetype).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create store type or name exists"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Store type created successfully", "storetype": storetype})
}

func GetStoretypes(c *gin.Context) {
	var storetypes []models.Storetype
	if err := config.DB.Find(&storetypes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve store types"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"storetypes": storetypes})
}

func GetStoretypeByID(c *gin.Context) {
	id := c.Param("id")
	var storetype models.Storetype

	if err := config.DB.First(&storetype, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Store type not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"storetype": storetype})
}

func UpdateStoretype(c *gin.Context) {
	id := c.Param("id")
	var storetype models.Storetype

	if err := config.DB.First(&storetype, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Store type not found"})
		return
	}

	var input struct {
		Name        string `json:"storetype_name"`
		Description string `json:"storetype_description"`
		Status      int    `json:"storetype_status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storetype.Name = input.Name
	storetype.Description = input.Description
	storetype.Status = input.Status

	if err := config.DB.Save(&storetype).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update store type"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Store type updated successfully", "storetype": storetype})
}

func DeleteStoretype(c *gin.Context) {
	id := c.Param("id")
	var storetype models.Storetype

	if err := config.DB.First(&storetype, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Store type not found"})
		return
	}

	if err := config.DB.Delete(&storetype).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete store type"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Store type deleted successfully"})
}