package controllers

import (
	"Eshop/models"
	"Eshop/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateLocation(c *gin.Context) {
	var input struct {
		Name    string `json:"location_name" binding:"required"`
		Country string `json:"location_country"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	location := models.Location{
		Name:    input.Name,
		Country: input.Country,
	}
	
	if err := config.DB.Create(&location).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create location or name exists"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Location created successfully", "location": location})
}

func GetLocations(c *gin.Context) {
	var locations []models.Location
	if err := config.DB.Find(&locations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve locations"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"locations": locations})
}

func GetLocationByID(c *gin.Context) {
	id := c.Param("id")
	var location models.Location

	if err := config.DB.First(&location, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Location not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"location": location})
}

func UpdateLocation(c *gin.Context) {
	id := c.Param("id")
	var location models.Location

	if err := config.DB.First(&location, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Location not found"})
		return
	}

	var input struct {
		Name    string `json:"location_name"`
		Country string `json:"location_country"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	location.Name = input.Name
	location.Country = input.Country

	if err := config.DB.Save(&location).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update location"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Location updated successfully", "location": location})
}


func DeleteLocation(c *gin.Context) {
	id := c.Param("id")
	var location models.Location

	if err := config.DB.First(&location, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Location not found"})
		return
	}

	if err := config.DB.Delete(&location).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete location"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Location deleted successfully"})
}


