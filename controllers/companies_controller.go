package controllers

import (
	"Eshop/models"
	"Eshop/config"
	"net/http"
	
	"github.com/gin-gonic/gin"
)

func CreateCompany(c *gin.Context) {
	var input struct {
		CompanyName string `json:"company_name" binding:"required"`
		Address     string `json:"address"`
		Phone       string `json:"phone"`
		Email       string `json:"email"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	company := models.Companies{
		CompanyName: input.CompanyName,
		Address:     input.Address,
		Phone:       input.Phone,
		Email:       input.Email,
	}

	if err := config.DB.Create(&company).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create company or name exists"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company created successfully", "company": company})
}

func GetCompanies(c *gin.Context) {
	var companies []models.Companies
	if err := config.DB.Find(&companies).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve companies"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"companies": companies})
}

func GetCompanyByID(c *gin.Context) {
	id := c.Param("id")
	var company models.Companies

	if err := config.DB.First(&company, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"company": company})
}

func UpdateCompany(c *gin.Context) {
	id := c.Param("id")
	var company models.Companies

	if err := config.DB.First(&company, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}

	var input struct {
		CompanyName string `json:"company_name"`
		Address     string `json:"address"`
		Phone       string `json:"phone"`
		Email       string `json:"email"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	company.CompanyName = input.CompanyName
	company.Address = input.Address
	company.Phone = input.Phone
	company.Email = input.Email

	if err := config.DB.Save(&company).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update company"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company updated successfully", "company": company})
}

// func DeleteCompany(c *gin.Context) {
// 	id := c.Param("id")
// 	var company models.Companies

// 	if err := config.DB.First(&company, id).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
// 		return
// 	}

// 	if err := config.DB.Delete(&company).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete company"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Company deleted successfully"})
// }




