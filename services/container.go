package services

import (
	"github.com/anggacipta/order-management-api/repositories"
	"gorm.io/gorm"
)

// ServiceContainer holds all services
type ServiceContainer struct {
	AuthService    AuthService
	ProductService ProductService
	OrderService   OrderService
}

// NewServiceContainer creates a new service container with all dependencies
func NewServiceContainer(db *gorm.DB) *ServiceContainer {
	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)
	productRepo := repositories.NewProductRepository(db)
	orderRepo := repositories.NewOrderRepository(db)

	// Initialize services
	authService := NewAuthService(userRepo)
	productService := NewProductService(productRepo)
	orderService := NewOrderService(orderRepo, productRepo, db)

	return &ServiceContainer{
		AuthService:    authService,
		ProductService: productService,
		OrderService:   orderService,
	}
}
