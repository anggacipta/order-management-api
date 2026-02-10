package routes

import (
	"github.com/anggacipta/order-management-api/controllers"
	"github.com/anggacipta/order-management-api/middlewares"
	"github.com/anggacipta/order-management-api/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, serviceContainer *services.ServiceContainer) {
	// Initialize controllers with dependency injection
	authController := controllers.NewAuthController(serviceContainer.AuthService)
	productController := controllers.NewProductController(serviceContainer.ProductService)
	orderController := controllers.NewOrderController(serviceContainer.OrderService)
	logController := controllers.NewLogController(serviceContainer.LogService)

	// Auth routes
	r.POST("/register", authController.Register)
	r.POST("/register-admin", authController.RegisterAdmin)
	r.POST("/login", authController.Login)

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
		auth.POST("/orders", orderController.CreateOrder)
		auth.GET("/orders", orderController.GetMyOrders)

		// Endpoint hanya untuk admin
		admin := auth.Group("/admin")
		admin.Use(middlewares.AdminOnly())
		{
			admin.GET("/dashboard", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "Welcome, admin!"})
			})
			// CRUD Produk
			admin.POST("/products", productController.CreateProduct)
			admin.GET("/products", productController.GetProducts)
			admin.GET("/products/:id", productController.GetProductByID)
			admin.PUT("/products/:id", productController.UpdateProduct)
			admin.DELETE("/products/:id", productController.DeleteProduct)

			// Paginated Products
			admin.GET("/products-paginated", productController.GetProductsPaginated)

			// Order endpoint log
			admin.GET("/logs", logController.GetAllLogs)
			admin.POST("/logs", logController.CreateLog)
		}
	}
}
