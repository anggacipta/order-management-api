package repositories

import (
	"github.com/anggacipta/order-management-api/models"
	"gorm.io/gorm"
)

type logRepository struct {
	db *gorm.DB // Uncomment and use if database operations are needed
}

// NewLogRepository creates a new log repository
func NewLogRepository(db *gorm.DB) LogRepository {
	return &logRepository{db: db}
}

func (r *logRepository) GetAll() ([]models.Log, error) {
	var logs []models.Log
	if err := r.db.Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}

func (r *logRepository) Create(log *models.Log) error {
	return r.db.Create(log).Error
}
