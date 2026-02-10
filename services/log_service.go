package services

import (
	"github.com/anggacipta/order-management-api/dto"
	"github.com/anggacipta/order-management-api/models"
	"github.com/anggacipta/order-management-api/repositories"
)

type logService struct {
	logRepo repositories.LogRepository
}

// NewLogService creates a new log service
func NewLogService(logRepo repositories.LogRepository) LogService {
	return &logService{logRepo: logRepo}
}

func (s *logService) GetAllLogs() ([]models.Log, error) {
	return s.logRepo.GetAll()
}

func (s *logService) CreateLog(req dto.LogRequest) error {
	log := &models.Log{
		Action:    req.Action,
		UserID:    req.UserID,
		Timestamp: req.Timestamp,
	}
	return s.logRepo.Create(log)
}
