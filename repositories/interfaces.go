package repositories

import (
	"github.com/anggacipta/order-management-api/models"
	"gorm.io/gorm"
)

// UserRepository interface untuk operasi database user
type UserRepository interface {
	Create(user *models.User) error
	GetByEmail(email string) (*models.User, error)
	GetByID(id uint) (*models.User, error)
}

// ProductRepository interface untuk operasi database product
type ProductRepository interface {
	GetByID(id uint) (*models.Product, error)
	GetByIDWithLock(tx *gorm.DB, id uint) (*models.Product, error)
	GetAllPaginated(page, limit int) ([]models.Product, int64, error)
	UpdateStock(tx *gorm.DB, product *models.Product) error
	GetAll() ([]models.Product, error)
	Create(product *models.Product) error
	Update(product *models.Product) error
	Delete(id uint) error
}

// OrderRepository interface untuk operasi database order
type OrderRepository interface {
	Create(tx *gorm.DB, order *models.Order) error
	GetByUserID(userID uint) ([]models.Order, error)
	GetByID(id uint) (*models.Order, error)
}

// LogRepository interface untuk operasi database log
type LogRepository interface {
	GetAll() ([]models.Log, error)
	Create(log *models.Log) error
}
