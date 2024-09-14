package controllers

import (
	"jobboard/models"
	"jobboard/utils"

	"github.com/gin-gonic/gin"
)

func RegisterStudent(c *gin.Context) {
	var student models.Student
	if err := c.BindJSON(&student); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	result := utils.DB.Create(&student)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(201, student)
}
