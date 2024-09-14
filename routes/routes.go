
package routes

import (
	"github.com/ishaandevburman/jobboard/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	student := router.Group("/student")
	{
		student.POST("/register", controllers.RegisterStudent)
		student.GET("/:registration_number", controllers.GetStudent)
		student.PUT("/:registration_number", controllers.UpdateStudent)
		student.DELETE("/:registration_number", controllers.DeleteStudent)
	}

	company := router.Group("/company")
	{
		company.POST("/add", controllers.AddCompany)
        company.GET("/:name", controllers.GetCompanyByName)
        company.PUT("/:name", controllers.UpdateCompanyByName)
        company.DELETE("/:name", controllers.DeleteCompanyByName)
        company.GET("/", controllers.ListCompanies)
		company.POST("/approve-job", controllers.ApproveJobListing)
	}

	job := router.Group("/job")
    {
        job.POST("/add", controllers.AddJob)          
        job.GET("/:id", controllers.GetJobByID)       
        job.PUT("/:id", controllers.UpdateJob)        
        job.DELETE("/:id", controllers.DeleteJob)     
        job.POST("/approve", controllers.ApproveJobListing) 
        job.GET("/", controllers.ListJobs)            
    }

}
