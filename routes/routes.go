package routes

import (
	"order-management-api/controllers"
	"order-management-api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Auth routes
	r.POST("/register", controllers.Register)
	r.POST("/register-admin", controllers.RegisterAdmin)
	r.POST("/login", controllers.Login)

	// Contoh: grup route yang butuh autentikasi
	auth := r.Group("/")
	auth.Use(middlewares.JWTAuthMiddleware())
	{
		// Endpoint yang hanya bisa diakses user login
		auth.GET("/me", func(c *gin.Context) {
			userID, _ := c.Get("user_id")
			role, _ := c.Get("role")
			c.JSON(200, gin.H{"user_id": userID, "role": role})
		})
		// Order endpoint (customer)
		auth.POST("/orders", controllers.CreateOrder)
		auth.GET("/orders", controllers.GetMyOrders)

		// Endpoint hanya untuk admin
		admin := auth.Group("/admin")
		admin.Use(middlewares.AdminOnly())
		{
			admin.GET("/dashboard", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Welcome, admin!"})
			})
			// CRUD Produk
			admin.POST("/products", controllers.CreateProduct)
			admin.GET("/products", controllers.GetProducts)
			admin.GET("/products/:id", controllers.GetProductByID)
			admin.PUT("/products/:id", controllers.UpdateProduct)
			admin.DELETE("/products/:id", controllers.DeleteProduct)
		}
	}
}
