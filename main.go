package main

import (
	"github.com/anggacipta/order-management-api/config"
	"github.com/anggacipta/order-management-api/models"
	"github.com/anggacipta/order-management-api/routes"
	"github.com/anggacipta/order-management-api/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Inisialisasi database
	models.ConnectDatabase()

	// Initialize service container with dependency injection
	serviceContainer := services.NewServiceContainer(models.DB)

	r := gin.Default()

	// Inisialisasi semua route dengan service container
	routes.SetupRoutes(r, serviceContainer)

	r.Run(":" + config.AppConfig.Port)
}
