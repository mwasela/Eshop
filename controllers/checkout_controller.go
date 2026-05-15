package controllers

import (
	"Eshop/config"
	"Eshop/models"
	"log"
	"net/http"
	
	"github.com/gin-gonic/gin"
)

func CreateCheckout(c *gin.Context) {
	var input struct {
		CartID     string  `json:"cart_id" binding:"required"`
		PaymentType *int    `json:"payment_type" binding:"required,oneof=1 2"`
		ClientFname string  `json:"client_fname" binding:"required"`
		ClientLname string  `json:"client_lname" binding:"required"`
		DeliveryCity string  `json:"delivery_city" binding:"required"`
		DeliveryStreetAddress string  `json:"delivery_street_address" binding:"required"`
		DeliveryCountry string  `json:"delivery_country" binding:"required"`
		DeliveryNotes string  `json:"delivery_notes"`
		ClientPhone string  `json:"client_phone"`
		TotalAmount *float64 `json:"total_amount" binding:"required"`
		PaymentFlag *int     `json:"payment_flag" binding:"omitempty,oneof=0 1"`
		CompanyID   *uint    `json:"company_id" binding:"required"`

	}

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("CreateCheckout bind error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid checkout payload", "details": err.Error()})
		return
	}

	paymentFlag := 1
	if input.PaymentFlag != nil {
		paymentFlag = *input.PaymentFlag
	}

	checkout := models.Checkout{
		CartID:     input.CartID,
		PaymentType: *input.PaymentType,
		ClientFname: input.ClientFname,
		ClientLname: input.ClientLname,
		DeliveryCity: input.DeliveryCity,
		DeliveryStreetAddress: input.DeliveryStreetAddress,
		DeliveryCountry: input.DeliveryCountry,
		DeliveryNotes: input.DeliveryNotes,
		ClientPhone: input.ClientPhone,
		TotalAmount: *input.TotalAmount,
		PaymentFlag: paymentFlag,
		CompanyID: *input.CompanyID,
	}
	
	if err := config.DB.Create(&checkout).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create checkout"})
		return
	}	

	c.JSON(http.StatusOK, gin.H{"message": "Checkout created successfully", "checkout": checkout})
}

func GetCheckouts(c *gin.Context) {
	var checkouts []models.Checkout
	if err := config.DB.Find(&checkouts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve checkouts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"checkouts": checkouts})
}

func GetCheckoutByID(c *gin.Context) {
	id := c.Param("id")
	var checkout models.Checkout

	if err := config.DB.First(&checkout, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Checkout not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"checkout": checkout})
}

func UpdateCheckout(c *gin.Context) {
	id := c.Param("id")
	var checkout models.Checkout

	if err := config.DB.First(&checkout, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Checkout not found"})
		return
	}

	var input struct {
		CartID     string  `json:"cart_id"`
		PaymentType int     `json:"payment_type" binding:"omitempty,oneof=1 2"`
		ClientFname string  `json:"client_fname"`
		ClientLname string  `json:"client_lname"`
		DeliveryCity string  `json:"delivery_city"`
		DeliveryStreetAddress string  `json:"delivery_street_address"`
		DeliveryCountry string  `json:"delivery_country"`
		DeliveryNotes string  `json:"delivery_notes"`
		ClientPhone string  `json:"client_phone"`
		TotalAmount float64 `json:"total_amount"`
		PaymentFlag int     `json:"payment_flag" binding:"omitempty,oneof=0 1"`
		CompanyID   uint    `json:"company_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	checkout.CartID = input.CartID
	checkout.PaymentType = input.PaymentType
	checkout.ClientFname = input.ClientFname
	checkout.ClientLname = input.ClientLname
	checkout.DeliveryCity = input.DeliveryCity
	checkout.DeliveryStreetAddress = input.DeliveryStreetAddress
	checkout.DeliveryCountry = input.DeliveryCountry
	checkout.DeliveryNotes = input.DeliveryNotes
	checkout.ClientPhone = input.ClientPhone
	checkout.TotalAmount = input.TotalAmount
	checkout.PaymentFlag = input.PaymentFlag
	checkout.CompanyID = input.CompanyID

	if err := config.DB.Save(&checkout).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update checkout"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Checkout updated successfully", "checkout": checkout})
}

func DeleteCheckout(c *gin.Context) {
	id := c.Param("id")
	var checkout models.Checkout

	if err := config.DB.First(&checkout, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Checkout not found"})
		return
	}

	if err := config.DB.Delete(&checkout).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete checkout"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Checkout deleted successfully"})
}

func GetCheckoutsByCartID(c *gin.Context) {
	cartID := c.Param("cart_id")
	var checkouts []models.Checkout

	if err := config.DB.Where("eshop_cart_id = ?", cartID).Find(&checkouts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve checkouts for cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"checkouts": checkouts})
}

