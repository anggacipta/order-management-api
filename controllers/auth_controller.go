package controllers

import (
	"github.com/anggacipta/order-management-api/dto"
	"github.com/anggacipta/order-management-api/helpers"
	"github.com/anggacipta/order-management-api/services"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// RegisterAdmin hanya untuk membuat user admin
func (ctrl *AuthController) RegisterAdmin(c *gin.Context) {
	var input dto.RegisterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.RespondValidationError(c, err)
		return
	}

	if err := ctrl.authService.RegisterAdmin(input); err != nil {
		helpers.RespondValidationError(c, err)
		return
	}

	c.JSON(200, gin.H{"message": "Registrasi admin berhasil"})
}

func (ctrl *AuthController) Register(c *gin.Context) {
	var input dto.RegisterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.RespondValidationError(c, err)
		return
	}

	if err := ctrl.authService.Register(input); err != nil {
		helpers.RespondValidationError(c, err)
		return
	}

	c.JSON(200, gin.H{"message": "Registrasi berhasil"})
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var input dto.LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.RespondValidationError(c, err)
		return
	}

	token, err := ctrl.authService.Login(input)
	if err != nil {
		helpers.RespondUnauthorized(c, err.Error())
		return
	}

	c.JSON(200, gin.H{"token": token})
}
