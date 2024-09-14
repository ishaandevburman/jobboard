package controllers

import (
	"jobboard/models"

	"github.com/gin-gonic/gin"
)

func ApproveJobListing(c *gin.Context) {
	var job models.Job
	if err := c.BindJSON(&job); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// Approve job logic...
	c.JSON(200, gin.H{"status": "Job listing approved"})
}
