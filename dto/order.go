package dto

type OrderItemRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,gte=1"`
}

type CreateOrderRequest struct {
	Items []OrderItemRequest `json:"items" binding:"required,dive,required"`
}
