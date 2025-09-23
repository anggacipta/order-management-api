package services

import (
	"errors"

	"github.com/anggacipta/order-management-api/dto"
	"github.com/anggacipta/order-management-api/models"
	"github.com/anggacipta/order-management-api/repositories"
	"gorm.io/gorm"
)

type orderService struct {
	orderRepo   repositories.OrderRepository
	productRepo repositories.ProductRepository
	db          *gorm.DB
}

// NewOrderService creates a new order service
func NewOrderService(orderRepo repositories.OrderRepository, productRepo repositories.ProductRepository, db *gorm.DB) OrderService {
	return &orderService{
		orderRepo:   orderRepo,
		productRepo: productRepo,
		db:          db,
	}
}

func (s *orderService) CreateOrder(userID uint, req dto.CreateOrderRequest) (*models.Order, error) {
	// Start database transaction
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	order := models.Order{UserID: userID}
	var orderItems []models.OrderItem

	// Process each item in transaction
	for _, item := range req.Items {
		// Lock the product row for update to prevent race condition
		product, err := s.productRepo.GetByIDWithLock(tx, item.ProductID)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		// Check stock availability
		if product.Stock < item.Quantity {
			tx.Rollback()
			return nil, errors.New("stok produk tidak cukup")
		}

		// Update stock
		product.Stock -= item.Quantity
		if err := s.productRepo.UpdateStock(tx, product); err != nil {
			tx.Rollback()
			return nil, err
		}

		orderItem := models.OrderItem{
			ProductID: product.ID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		}
		orderItems = append(orderItems, orderItem)
	}

	order.OrderItems = orderItems

	// Create order in transaction
	if err := s.orderRepo.Create(tx, &order); err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &order, nil
}

func (s *orderService) GetMyOrders(userID uint) ([]models.Order, error) {
	return s.orderRepo.GetByUserID(userID)
}

func (s *orderService) GetOrderByID(id uint) (*models.Order, error) {
	return s.orderRepo.GetByID(id)
}
