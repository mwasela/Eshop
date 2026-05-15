package controllers

import (
	"Eshop/models"
	"Eshop/config"
	"net/http"
	
	"github.com/gin-gonic/gin"
)


func CreateCart(c *gin.Context) {
	var input struct {
		CartID     string  `json:"cart_id" binding:"required"`
		ProductID  uint    `json:"product_id" binding:"required"`
		Quantity   int     `json:"quantity" binding:"required"`
		UnitPrice  float64 `json:"unit_price" binding:"required"`
		UserID     *uint   `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Calculate TotalPrice
	totalPrice := input.UnitPrice * float64(input.Quantity)

	cart := models.Cart{
		CartID:     input.CartID,
		ProductID:  input.ProductID,
		Quantity:   input.Quantity,
		UnitPrice:  input.UnitPrice,
		TotalPrice: totalPrice,
		UserID:     input.UserID,
	}
	
	if err := config.DB.Create(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart created successfully", "cart": cart})
}

//insert multipke items to cart with same cart_id
func CreateMultipleCartItems(c *gin.Context) {
	var input []struct {
		CartID     string  `json:"cart_id" binding:"required"`
		ProductID  uint    `json:"product_id" binding:"required"`
		Quantity   int     `json:"quantity" binding:"required"`
		UnitPrice  float64 `json:"unit_price" binding:"required"`
		UserID     *uint   `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var carts []models.Cart
	for _, item := range input {
		// Calculate TotalPrice
		totalPrice := item.UnitPrice * float64(item.Quantity)

		cart := models.Cart{
			CartID:     item.CartID,
			ProductID:  item.ProductID,
			Quantity:   item.Quantity,
			UnitPrice:  item.UnitPrice,
			TotalPrice: totalPrice,
			UserID:     item.UserID,
		}
		carts = append(carts, cart)
	}

	if err := config.DB.Create(&carts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cart items"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart items created successfully", "cart_items": carts})
}

func GetCartItemsByCartID(c *gin.Context) {
	cartID := c.Param("cart_id")
	var cartItems []models.Cart

	if err := config.DB.Where("eshop_cart_id = ?", cartID).Preload("Product").Preload("User").Find(&cartItems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cart items"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"cart_items": cartItems})
}

func GetCarts(c *gin.Context) {
	var carts []models.Cart
	if err := config.DB.Preload("Product").Preload("User").Find(&carts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve carts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"carts": carts})
}

func GetCartByID(c *gin.Context) {
	cartID := c.Param("cart_id")
	var cart models.Cart

	if err := config.DB.Preload("Product").Preload("User").First(&cart, "eshop_cart_id = ?", cartID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cart": cart})
}

func UpdateCart(c *gin.Context) {
	cartID := c.Param("cart_id")
	var cart models.Cart

	if err := config.DB.First(&cart, "eshop_cart_id = ?", cartID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}

	var input struct {
		CartID     string  `json:"cart_id"`
		ProductID  uint    `json:"product_id"`
		Quantity   int     `json:"quantity"`
		UnitPrice  float64 `json:"unit_price"`
		TotalPrice float64 `json:"total_price"`
		UserID     *uint   `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cart.CartID = input.CartID

	if input.ProductID != 0 {
		cart.ProductID = input.ProductID
	}
	if input.Quantity != 0 {
		cart.Quantity = input.Quantity
	}
	if input.UnitPrice != 0 {
		cart.UnitPrice = input.UnitPrice
	}
	if input.TotalPrice != 0 {
		cart.TotalPrice = input.TotalPrice
	}
	if input.UserID != nil {
		cart.UserID = input.UserID
	}

	if err := config.DB.Save(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart updated successfully", "cart": cart})
}

func DeleteCart(c *gin.Context) {
	cartID := c.Param("cart_id")
	var cart models.Cart

	if err := config.DB.First(&cart, "eshop_cart_id = ?", cartID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}

	if err := config.DB.Delete(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart deleted successfully"})
}


