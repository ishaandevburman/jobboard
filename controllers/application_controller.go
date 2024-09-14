package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ishaandevburman/jobboard/models"
	"github.com/ishaandevburman/jobboard/utils"
)

func ApplyForJob(c *gin.Context) {
	var application models.Application
	if err := c.BindJSON(&application); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := utils.DB.Create(&application)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, application)
}

func GetApplication(c *gin.Context) {
	id := c.Param("id")
	var application models.Application
	result := utils.DB.First(&application, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
		return
	}
	c.JSON(http.StatusOK, application)
}

func UpdateApplication(c *gin.Context) {
	id := c.Param("id")
	var application models.Application
	result := utils.DB.First(&application, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
		return
	}

	if err := c.BindJSON(&application); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	utils.DB.Save(&application)
	c.JSON(http.StatusOK, application)
}

func DeleteApplication(c *gin.Context) {
	id := c.Param("id")
	result := utils.DB.Delete(&models.Application{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Application deleted"})
}

func ListApplications(c *gin.Context) {
	var applications []models.Application
	result := utils.DB.Find(&applications)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, applications)
}
