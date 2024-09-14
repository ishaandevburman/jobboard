package controllers

import (
	"github.com/ishaandevburman/jobboard/models"
	"github.com/ishaandevburman/jobboard/utils"

	"github.com/gin-gonic/gin"
)

func AddCompany(c *gin.Context) {
	var company models.Company
	if err := c.BindJSON(&company); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result := utils.DB.Create(&company)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(201, company)
}
