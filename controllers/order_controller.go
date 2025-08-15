package controllers

import (
	"errors"
	"order-management-api/dto"
	"order-management-api/helpers"
	"order-management-api/models"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var input dto.CreateOrderRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.RespondValidationError(c, err)
		return
	}
	userID, _ := c.Get("user_id")
	order := models.Order{UserID: userID.(uint)}
	var orderItems []models.OrderItem
	for _, item := range input.Items {
		var product models.Product
		if err := models.DB.First(&product, item.ProductID).Error; err != nil {
			helpers.RespondNotFound(c, "Produk tidak ditemukan")
			return
		}
		if product.Stock < item.Quantity {
			helpers.RespondValidationError(c, errors.New("stok produk tidak cukup"))
			return
		}
		product.Stock -= item.Quantity
		models.DB.Save(&product)
		orderItem := models.OrderItem{
			ProductID: product.ID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		}
		orderItems = append(orderItems, orderItem)
	}
	order.OrderItems = orderItems
	if err := models.DB.Create(&order).Error; err != nil {
		helpers.RespondInternalError(c, err)
		return
	}
	c.JSON(200, order)
}

func GetMyOrders(c *gin.Context) {
	userID, _ := c.Get("user_id")
	var orders []models.Order
	models.DB.Preload("OrderItems.Product").Where("user_id = ?", userID).Find(&orders)
	c.JSON(200, orders)
}
