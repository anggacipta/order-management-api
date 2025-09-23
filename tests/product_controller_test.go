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

func TestMain(m *testing.M) {
	// Setup DB in-memory untuk test
	models.SetupTestDB()
	m.Run()
}

func setupTestRouter() *gin.Engine {
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

	// Initialize product controller
	productController := controllers.NewProductController(serviceContainer.ProductService)

	r := gin.Default()
	r.POST("/admin/products", productController.CreateProduct)
	return r
}

func TestCreateProduct_Success(t *testing.T) {
	r := setupTestRouter()
	input := dto.CreateProductRequest{
		Name:        "Test Product",
		Description: "Test Desc",
		Price:       100,
		Stock:       10,
	}
	jsonValue, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/admin/products", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestCreateProduct_ValidationError(t *testing.T) {
	r := setupTestRouter()
	input := dto.CreateProductRequest{
		Name:        "",
		Description: "",
		Price:       -1,
		Stock:       -1,
	}
	jsonValue, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/admin/products", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}
