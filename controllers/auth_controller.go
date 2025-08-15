package controllers

import (
	"errors"
	"strings"

	"github.com/anggacipta/order-management-api/dto"
	"github.com/anggacipta/order-management-api/helpers"
	"github.com/anggacipta/order-management-api/models"
	"github.com/anggacipta/order-management-api/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// RegisterAdmin hanya untuk membuat user admin
func RegisterAdmin(c *gin.Context) {
	var input dto.RegisterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.RespondValidationError(c, err)
		return
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	user := models.User{Name: input.Name, Email: input.Email, Password: string(hashedPassword), Role: "admin"}
	if err := models.DB.Create(&user).Error; err != nil {
		if strings.Contains(err.Error(), "UNIQUE") && strings.Contains(err.Error(), "users.email") {
			helpers.RespondValidationError(c, errors.New("email sudah terdaftar"))
			return
		}
		helpers.RespondValidationError(c, err)
		return
	}
	c.JSON(200, gin.H{"message": "Registrasi admin berhasil"})
}

func Register(c *gin.Context) {
	var input dto.RegisterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.RespondValidationError(c, err)
		return
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	user := models.User{Name: input.Name, Email: input.Email, Password: string(hashedPassword), Role: "customer"}
	if err := models.DB.Create(&user).Error; err != nil {
		if strings.Contains(err.Error(), "UNIQUE") && strings.Contains(err.Error(), "users.email") {
			helpers.RespondValidationError(c, errors.New("email sudah terdaftar"))
			return
		}
		helpers.RespondValidationError(c, err)
		return
	}
	c.JSON(200, gin.H{"message": "Registrasi berhasil"})
}

func Login(c *gin.Context) {
	var input dto.LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.RespondValidationError(c, err)
		return
	}
	var user models.User
	if err := models.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		helpers.RespondUnauthorized(c, "Email tidak ditemukan")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		helpers.RespondUnauthorized(c, "Password salah")
		return
	}
	token, err := utils.GenerateJWT(user)
	if err != nil {
		helpers.RespondInternalError(c, err)
		return
	}
	c.JSON(200, gin.H{"token": token})
}
