package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ishaandevburman/jobboard/models"
	"github.com/ishaandevburman/jobboard/utils"
)

func RegisterStudent(c *gin.Context) {
	var student models.Student
	if err := c.BindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingStudent models.Student
	result := utils.DB.Where("registration_number = ?", student.RegistrationNumber).First(&existingStudent)
	if result.Error == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Registration number already exists"})
		return
	}

	result = utils.DB.Create(&student)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, student)
}

func GetStudent(c *gin.Context) {
	registrationNumber := c.Param("registration_number")
	var student models.Student
	result := utils.DB.Where("registration_number = ?", registrationNumber).First(&student)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

func UpdateStudent(c *gin.Context) {
	registrationNumber := c.Param("registration_number")
	var student models.Student
	result := utils.DB.Where("registration_number = ?", registrationNumber).First(&student)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	if err := c.BindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	utils.DB.Save(&student)
	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	registrationNumber := c.Param("registration_number")
	result := utils.DB.Where("registration_number = ?", registrationNumber).Delete(&models.Student{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Student deleted"})
}
