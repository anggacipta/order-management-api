package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"order-management-api/controllers"
	"order-management-api/dto"
	"order-management-api/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupAuthTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	return r
}

func TestRegister_Success(t *testing.T) {
	models.SetupTestDB() // Pastikan DB test di-reset
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
	models.SetupTestDB()
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
	models.SetupTestDB()
	r := setupAuthTestRouter()
	// Register user dulu
	registerInput := dto.RegisterRequest{
		Name:     "Login User",
		Email:    "loginuser@example.com",
		Password: "password123",
	}
	jsonValue, _ := json.Marshal(registerInput)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Login
	loginInput := dto.LoginRequest{
		Email:    "loginuser@example.com",
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
	models.SetupTestDB()
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
