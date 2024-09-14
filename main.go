// main.go (for migration)
package main

import (
	"github.com/ishaandevburman/jobboard/models"
	"github.com/ishaandevburman/jobboard/routes"
	"github.com/ishaandevburman/jobboard/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.ConnectDatabase()

	// Auto-migrate the models
	utils.DB.AutoMigrate(&models.Student{}, &models.Company{}, &models.Job{}, &models.Application{})

	router := gin.Default()
	routes.SetupRoutes(router)

	router.Run(":8080")
}
