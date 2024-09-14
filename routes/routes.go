// routes/routes.go
package routes

import (
	"github.com/ishaandevburman/jobboard/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	student := router.Group("/student")
	{
		student.POST("/register", controllers.RegisterStudent)
		// other student routes...
	}

	company := router.Group("/company")
	{
		company.POST("/add", controllers.AddCompany)
		// other company routes...
	}

	placementCell := router.Group("/placement")
	{
		placementCell.POST("/approve-job", controllers.ApproveJobListing)
		// other placement cell routes...
	}
}
