package controllers

import (
	"net/http"
	"strconv"

	"github.com/anggacipta/order-management-api/dto"
	"github.com/anggacipta/order-management-api/helpers"
	"github.com/anggacipta/order-management-api/services"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (ctrl *ProductController) CreateProduct(c *gin.Context) {
	var input dto.CreateProductRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.RespondValidationError(c, err)
		return
	}

	product, err := ctrl.productService.Create(input)
	if err != nil {
		helpers.RespondInternalError(c, err)
		return
	}

	c.JSON(200, product)
}

func (ctrl *ProductController) GetProducts(c *gin.Context) {
	products, err := ctrl.productService.GetAll()
	if err != nil {
		helpers.RespondInternalError(c, err)
		return
	}

	c.JSON(http.StatusOK, products)
}

func (ctrl *ProductController) GetProductsPaginated(c *gin.Context) {
	var req dto.PaginationRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		helpers.RespondValidationError(c, err)
		return
	}

	result, err := ctrl.productService.GetAllPaginated(req.Page, req.Limit)
	if err != nil {
		helpers.RespondInternalError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (ctrl *ProductController) GetProductByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		helpers.RespondValidationError(c, err)
		return
	}

	product, err := ctrl.productService.GetByID(uint(id))
	if err != nil {
		helpers.RespondNotFound(c, err.Error())
		return
	}

	c.JSON(200, product)
}

func (ctrl *ProductController) UpdateProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		helpers.RespondValidationError(c, err)
		return
	}

	var input dto.UpdateProductRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		helpers.RespondValidationError(c, err)
		return
	}

	product, err := ctrl.productService.Update(uint(id), input)
	if err != nil {
		helpers.RespondNotFound(c, err.Error())
		return
	}

	c.JSON(200, product)
}

func (ctrl *ProductController) DeleteProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		helpers.RespondValidationError(c, err)
		return
	}

	if err := ctrl.productService.Delete(uint(id)); err != nil {
		helpers.RespondNotFound(c, err.Error())
		return
	}

	c.JSON(200, gin.H{"message": "Produk berhasil dihapus"})
}
