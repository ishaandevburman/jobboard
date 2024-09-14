package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ishaandevburman/jobboard/models"
	"github.com/ishaandevburman/jobboard/utils"
)

func AddCompany(c *gin.Context) {
	var company models.Company
	if err := c.BindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingCompany models.Company
	result := utils.DB.Where("name = ?", company.Name).First(&existingCompany)
	if result.Error == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Company name already exists"})
		return
	}

	result = utils.DB.Create(&company)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, company)
}

func GetCompanyByName(c *gin.Context) {
	name := c.Param("name")
	var company models.Company
	result := utils.DB.Where("name = ?", name).First(&company)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}
	c.JSON(http.StatusOK, company)
}

func UpdateCompanyByName(c *gin.Context) {
	name := c.Param("name")
	var company models.Company
	result := utils.DB.Where("name = ?", name).First(&company)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}

	if err := c.BindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	utils.DB.Save(&company)
	c.JSON(http.StatusOK, company)
}

func DeleteCompanyByName(c *gin.Context) {
	name := c.Param("name")
	result := utils.DB.Where("name = ?", name).Delete(&models.Company{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Company deleted"})
}

func ListCompanies(c *gin.Context) {
	var companies []models.Company
	result := utils.DB.Find(&companies)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, companies)
}
