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

func setupAuthTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	// Setup test database
	models.SetupTestDB()

	// Setup test config BEFORE creating service container
	config.AppConfig = &config.Config{
		JWTSecret: "test-secret-key-for-testing",
		Port:      "8080",
		DBPath:    ":memory:",
		AppEnv:    "test",
	}

	// Initialize service container
	serviceContainer := services.NewServiceContainer(models.DB)

	// Initialize auth controller
	authController := controllers.NewAuthController(serviceContainer.AuthService)

	r := gin.Default()
	r.POST("/register", authController.Register)
	r.POST("/login", authController.Login)
	return r
}

func TestRegister_Success(t *testing.T) {
	r := setupAuthTestRouter()
	input := dto.RegisterRequest{
		Name:     "Test User",
		Email:    "testuser@example.com",
		Password: "password123",
	}
	jsonValue, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestRegister_ValidationError(t *testing.T) {
	r := setupAuthTestRouter()
	input := dto.RegisterRequest{
		Name:     "",
		Email:    "",
		Password: "",
	}
	jsonValue, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 400, w.Code)
}

func TestLogin_Success(t *testing.T) {
	r := setupAuthTestRouter()
	// Register user dulu
	registerInput := dto.RegisterRequest{
		Name:     "Login User",
		Email:    "testuser@example.com",
		Password: "password123",
	}
	jsonValue, _ := json.Marshal(registerInput)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Login
	loginInput := dto.LoginRequest{
		Email:    "testuser@example.com",
		Password: "password123",
	}
	loginJson, _ := json.Marshal(loginInput)
	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(loginJson))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestLogin_InvalidCredentials(t *testing.T) {
	r := setupAuthTestRouter()
	loginInput := dto.LoginRequest{
		Email:    "notfound@example.com",
		Password: "wrongpassword",
	}
	loginJson, _ := json.Marshal(loginInput)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(loginJson))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Code)
}
