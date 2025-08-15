package main

import (
	"github.com/anggacipta/order-management-api/models"
	"github.com/anggacipta/order-management-api/routes"

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
