package controllers

import (
	"net/http"
	"order-management-api/dto"
	"order-management-api/helpers"
	"order-management-api/models"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var input dto.ProductRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.RespondValidationError(c, err)
		return
	}
	product := models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
	}
	if err := models.DB.Create(&product).Error; err != nil {
		helpers.RespondInternalError(c, err)
		return
	}
	c.JSON(200, product)
}

func GetProducts(c *gin.Context) {
	var products []models.Product
	models.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

func GetProductByID(c *gin.Context) {
	var product models.Product
	id := c.Param("id")
	if err := models.DB.First(&product, id).Error; err != nil {
		helpers.RespondNotFound(c, "Produk tidak ditemukan")
		return
	}
	c.JSON(200, product)
}

func UpdateProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")
	if err := models.DB.First(&product, id).Error; err != nil {
		helpers.RespondNotFound(c, "Produk tidak ditemukan")
		return
	}
	var input dto.ProductRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.RespondValidationError(c, err)
		return
	}
	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	product.Stock = input.Stock
	models.DB.Save(&product)
	c.JSON(200, product)
}

func DeleteProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")
	if err := models.DB.First(&product, id).Error; err != nil {
		helpers.RespondNotFound(c, "Produk tidak ditemukan")
		return
	}
	models.DB.Delete(&product)
	c.JSON(200, gin.H{"message": "Produk berhasil dihapus"})
}
