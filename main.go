package main

import (
	"order-management-api/models"
	"order-management-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi database
	models.ConnectDatabase()

	r := gin.Default()

	// Inisialisasi semua route
	routes.SetupRoutes(r)

	r.Run(":8080")
}
