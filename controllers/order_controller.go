package controllers

import (
	"github.com/anggacipta/order-management-api/dto"
	"github.com/anggacipta/order-management-api/helpers"
	"github.com/anggacipta/order-management-api/services"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService services.OrderService
}

func NewOrderController(orderService services.OrderService) *OrderController {
	return &OrderController{orderService: orderService}
}

func (ctrl *OrderController) CreateOrder(c *gin.Context) {
	var input dto.CreateOrderRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.RespondValidationError(c, err)
		return
	}

	userID, _ := c.Get("user_id")

	order, err := ctrl.orderService.CreateOrder(userID.(uint), input)
	if err != nil {
		helpers.RespondValidationError(c, err)
		return
	}

	c.JSON(200, order)
}

func (ctrl *OrderController) GetMyOrders(c *gin.Context) {
	userID, _ := c.Get("user_id")

	orders, err := ctrl.orderService.GetMyOrders(userID.(uint))
	if err != nil {
		helpers.RespondInternalError(c, err)
		return
	}

	c.JSON(200, orders)
}
