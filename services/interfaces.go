package services

import (
	"github.com/anggacipta/order-management-api/dto"
	"github.com/anggacipta/order-management-api/models"
)

// AuthService interface untuk business logic authentication
type AuthService interface {
	RegisterAdmin(req dto.RegisterRequest) error
	Register(req dto.RegisterRequest) error
	Login(req dto.LoginRequest) (string, error)
}

// ProductService interface untuk business logic product
type ProductService interface {
	GetAll() ([]models.Product, error)
	GetAllPaginated(page, limit int) (*dto.PaginationResponse, error)
	GetByID(id uint) (*models.Product, error)
	Create(req dto.CreateProductRequest) (*models.Product, error)
	Update(id uint, req dto.UpdateProductRequest) (*models.Product, error)
	Delete(id uint) error
}

// OrderService interface untuk business logic order
type OrderService interface {
	CreateOrder(userID uint, req dto.CreateOrderRequest) (*models.Order, error)
	GetMyOrders(userID uint) ([]models.Order, error)
	GetOrderByID(id uint) (*models.Order, error)
}

// LogService interface untuk business logic log
type LogService interface {
	GetAllLogs() ([]models.Log, error)
	CreateLog(req dto.LogRequest) error
}
