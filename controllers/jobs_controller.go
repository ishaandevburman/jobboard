package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ishaandevburman/jobboard/models"
	"github.com/ishaandevburman/jobboard/utils"
)

func AddJob(c *gin.Context) {
	var job models.Job
	if err := c.BindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := utils.DB.Create(&job)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, job)
}

func GetJobByID(c *gin.Context) {
	id := c.Param("id")
	var job models.Job
	result := utils.DB.First(&job, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
		return
	}
	c.JSON(http.StatusOK, job)
}

func UpdateJob(c *gin.Context) {
	id := c.Param("id")
	var job models.Job
	result := utils.DB.First(&job, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
		return
	}

	if err := c.BindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	utils.DB.Save(&job)
	c.JSON(http.StatusOK, job)
}

func DeleteJob(c *gin.Context) {
	id := c.Param("id")
	result := utils.DB.Delete(&models.Job{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Job deleted"})
}

func ListJobs(c *gin.Context) {
	var jobs []models.Job
	result := utils.DB.Find(&jobs)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, jobs)
}

func ApproveJobListing(c *gin.Context) {
	var job models.Job
	if err := c.BindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := utils.DB.Model(&models.Job{}).Where("id = ?", job.ID).Update("status", "approved")
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Job listing approved"})
}
