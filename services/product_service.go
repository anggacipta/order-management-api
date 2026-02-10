package services

import (
	"github.com/anggacipta/order-management-api/dto"
	"github.com/anggacipta/order-management-api/models"
	"github.com/anggacipta/order-management-api/repositories"
)

type productService struct {
	productRepo repositories.ProductRepository
}

// NewProductService creates a new product service
func NewProductService(productRepo repositories.ProductRepository) ProductService {
	return &productService{productRepo: productRepo}
}

func (s *productService) GetAll() ([]models.Product, error) {
	return s.productRepo.GetAll()
}

func (s *productService) GetAllPaginated(page, limit int) (*dto.PaginationResponse, error) {
	// Set default values
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	products, totalRows, err := s.productRepo.GetAllPaginated(page, limit)
	if err != nil {
		return nil, err
	}

	totalPages := int(totalRows) / limit
	if int(totalRows)%limit > 0 {
		totalPages++
	}

	return &dto.PaginationResponse{
		Data:       products,
		Page:       page,
		Limit:      limit,
		TotalRows:  totalRows,
		TotalPages: totalPages,
	}, nil
}

func (s *productService) GetByID(id uint) (*models.Product, error) {
	return s.productRepo.GetByID(id)
}

func (s *productService) Create(req dto.CreateProductRequest) (*models.Product, error) {
	product := models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}

	if err := s.productRepo.Create(&product); err != nil {
		return nil, err
	}

	return &product, nil
}

func (s *productService) Update(id uint, req dto.UpdateProductRequest) (*models.Product, error) {
	product, err := s.productRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.Stock = req.Stock

	if err := s.productRepo.Update(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *productService) Delete(id uint) error {
	// Check if product exists first
	_, err := s.productRepo.GetByID(id)
	if err != nil {
		return err
	}

	return s.productRepo.Delete(id)
}
