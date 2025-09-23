package controllers

import (
	"errors"

	"github.com/anggacipta/order-management-api/dto"
	"github.com/anggacipta/order-management-api/helpers"
	"github.com/anggacipta/order-management-api/models"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var input dto.CreateOrderRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.RespondValidationError(c, err)
		return
	}

	userID, _ := c.Get("user_id")

	// Start database transaction
	tx := models.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		helpers.RespondInternalError(c, err)
		return
	}

	order := models.Order{UserID: userID.(uint)}
	var orderItems []models.OrderItem

	// Process each item in transaction
	for _, item := range input.Items {
		var product models.Product
		// Lock the product row for update to prevent race condition
		if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&product, item.ProductID).Error; err != nil {
			tx.Rollback()
			helpers.RespondNotFound(c, "Produk tidak ditemukan")
			return
		}

		// Check stock availability
		if product.Stock < item.Quantity {
			tx.Rollback()
			helpers.RespondValidationError(c, errors.New("stok produk tidak cukup"))
			return
		}

		// Update stock
		product.Stock -= item.Quantity
		if err := tx.Save(&product).Error; err != nil {
			tx.Rollback()
			helpers.RespondInternalError(c, err)
			return
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
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		helpers.RespondInternalError(c, err)
		return
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
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
