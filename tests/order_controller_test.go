package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anggacipta/order-management-api/config"
	"github.com/anggacipta/order-management-api/controllers"
	"github.com/anggacipta/order-management-api/dto"
	"github.com/anggacipta/order-management-api/models"
	"github.com/anggacipta/order-management-api/services"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupOrderTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	models.SetupTestDB()

	// Setup test config
	config.AppConfig = &config.Config{
		JWTSecret: "test-secret-key-for-testing",
		Port:      "8080",
		DBPath:    ":memory:",
		AppEnv:    "test",
	}

	// Initialize service container
	serviceContainer := services.NewServiceContainer(models.DB)

	// Initialize order controller
	orderController := controllers.NewOrderController(serviceContainer.OrderService)

	r := gin.Default()
	// Middleware inject user_id
	r.Use(func(c *gin.Context) { c.Set("user_id", uint(1)) })
	r.POST("/orders", orderController.CreateOrder)
	r.GET("/orders", orderController.GetMyOrders)
	return r
}

func TestCreateOrder_Success(t *testing.T) {
	r := setupOrderTestRouter()
	// Insert product ke DB
	product := models.Product{Name: "Order Product", Description: "desc", Price: 100, Stock: 10}
	models.DB.Create(&product)
	// Simulasi user login (inject user_id ke context)
	orderInput := dto.CreateOrderRequest{
		Items: []dto.OrderItemRequest{{ProductID: product.ID, Quantity: 2}},
	}
	jsonValue, _ := json.Marshal(orderInput)
	req, _ := http.NewRequest("POST", "/orders", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestCreateOrder_StokKurang(t *testing.T) {
	r := setupOrderTestRouter()
	product := models.Product{Name: "Order Product", Description: "desc", Price: 100, Stock: 1}
	models.DB.Create(&product)
	orderInput := dto.CreateOrderRequest{
		Items: []dto.OrderItemRequest{{ProductID: product.ID, Quantity: 2}},
	}
	jsonValue, _ := json.Marshal(orderInput)
	req, _ := http.NewRequest("POST", "/orders", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}
