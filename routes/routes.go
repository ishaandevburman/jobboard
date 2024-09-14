package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ishaandevburman/jobboard/controllers"
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

	application := router.Group("/application")
	{
		application.POST("/apply", controllers.ApplyForJob)
		application.GET("/:id", controllers.GetApplication)
		application.PUT("/:id", controllers.UpdateApplication)
		application.DELETE("/:id", controllers.DeleteApplication)
		application.GET("/", controllers.ListApplications)
	}
}
