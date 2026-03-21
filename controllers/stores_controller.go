package controllers

import(
	"Eshop/models"
	"Eshop/config"
	"net/http"

	"github.com/gin-gonic/gin"
)


func CreateStore(c *gin.Context) {
	var input struct {
		StoreName string `json:"store_name" binding:"required"`
		LocationID uint   `json:"location_id"`
		Area	  string `json:"area"`
		CompanyID uint   `json:"company_id"`
		StoreTypeID uint   `json:"store_type_id"`
		Status    int    `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	store := models.Stores{
		StoreName: input.StoreName,
		LocationID: input.LocationID,
		Area: input.Area,
		CompanyID: input.CompanyID,
		StoretypeID: input.StoreTypeID,
		Status: input.Status,
	}

	if err := config.DB.Create(&store).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create store or name exists"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Store created successfully", "store": store})
}

func GetStores(c *gin.Context) {
	       var stores []models.Stores
	       if err := config.DB.Preload("Location").Preload("Companies").Preload("Storetype").Find(&stores).Error; err != nil {
		       c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve stores"})
		       return
	       }

	       c.JSON(http.StatusOK, gin.H{"stores": stores})
}

func GetStoreByID(c *gin.Context) {
	       id := c.Param("id")
	       var store models.Stores

	       if err := config.DB.Preload("Location").Preload("Companies").Preload("Storetype").First(&store, id).Error; err != nil {
		       c.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
		       return
	       }

	       c.JSON(http.StatusOK, gin.H{"store": store})
}

func UpdateStore(c *gin.Context) {
	id := c.Param("id")
	var store models.Stores

	if err := config.DB.First(&store, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
		return
	}

	var input struct {
		StoreName string `json:"store_name"`
		LocationID uint   `json:"location_id"`
		Area	  string `json:"area"`
		CompanyID uint   `json:"company_id"`
		StoreTypeID uint   `json:"store_type_id"`
		Status    int    `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	store.StoreName = input.StoreName
	store.LocationID = input.LocationID
	store.Area = input.Area
	store.CompanyID = input.CompanyID
	store.StoretypeID = input.StoreTypeID	
	store.Status = input.Status

	if err := config.DB.Save(&store).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update store"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Store updated successfully", "store": store})
}

func DeleteStore(c *gin.Context) {
	id := c.Param("id")
	var store models.Stores

	if err := config.DB.First(&store, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
		return
	}

	if err := config.DB.Delete(&store).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete store"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Store deleted successfully"})
}

