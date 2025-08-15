package dto

type ProductRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int    `json:"price" binding:"required,gte=0"`
	Stock       int    `json:"stock" binding:"required,gte=0"`
}
