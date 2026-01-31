package controllers

import(
	"Eshop/models"
	"Eshop/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var input struct {
		Name        string  `json:"name" binding:"required"`
		SKU		 string  `json:"sku" binding:"required"`
		OEM 	  string  `json:"oem"`
		Description string  `json:"description"`
		CategoryID  uint    `json:"category_id"`
	}
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	product := models.Product{
		Name:        input.Name,
		SKU:	     input.SKU,
		OEM:	     input.OEM,
		Description: input.Description,
		CategoryID:  input.CategoryID,
	}

	if err := config.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product or SKU exists"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully", "product": product})
}

func GetProducts(c *gin.Context) {
	var products []models.Product
	if err := config.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var input struct {
		Name        string  `json:"name"`
		SKU		 string  `json:"sku"`
		OEM 	  string  `json:"oem"`
		Description string  `json:"description"`
		CategoryID  uint    `json:"category_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProduct := models.Product{
		Name:        input.Name,
		SKU:	     input.SKU,
		OEM:	     input.OEM,
		Description: input.Description,
		CategoryID:  input.CategoryID,
	}

	if err := config.DB.Model(&product).Updates(updatedProduct).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully", "product": product})
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if err := config.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}