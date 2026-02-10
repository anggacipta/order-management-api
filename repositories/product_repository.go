package repositories

import (
	"errors"

	"github.com/anggacipta/order-management-api/models"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

// NewProductRepository creates a new product repository
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetByID(id uint) (*models.Product, error) {
	var product models.Product
	if err := r.db.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("produk tidak ditemukan")
		}
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) GetAllPaginated(page, limit int) ([]models.Product, int64, error) {
	var products []models.Product
	var totalRows int64

	offset := (page - 1) * limit

	if err := r.db.Model(&models.Product{}).Count(&totalRows).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, totalRows, nil
}

func (r *productRepository) GetByIDWithLock(tx *gorm.DB, id uint) (*models.Product, error) {
	var product models.Product
	if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("produk tidak ditemukan")
		}
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) UpdateStock(tx *gorm.DB, product *models.Product) error {
	return tx.Save(product).Error
}

func (r *productRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) Update(product *models.Product) error {
	// Start transaction for race condition prevention
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// Lock the product row for update
	var existingProduct models.Product
	if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&existingProduct, product.ID).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("produk tidak ditemukan")
		}
		return err
	}

	// Update the product
	if err := tx.Save(product).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit transaction
	return tx.Commit().Error
}

func (r *productRepository) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}
