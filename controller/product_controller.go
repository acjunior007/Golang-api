package controller

import (
	"net/http"
	"strconv"

	"go-api/model"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	// usecase.ProductUsecase
	productUseCase usecase.ProductUsecase
}

func NewProductController(useCase usecase.ProductUsecase) ProductController {
	return ProductController{
		productUseCase: useCase,
	}
}

func (pc *ProductController) GetProducts(ctx *gin.Context) {

	products, err := pc.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (pc *ProductController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	createdProduct, err := pc.productUseCase.CreateProduct(product)
	if err != nil {
		if err == usecase.ErrInvalidProductData {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product data"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	ctx.JSON(http.StatusCreated, createdProduct)
}

func (pc *ProductController) GetProductByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		response := model.Response{Message: "Invalid ID - Null or empty"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productID, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{Message: "Invalid ID - Not a number"}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	product, err := pc.productUseCase.GetProductByID(productID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if product == nil {
		response := model.Response{Message: "Product not found"}
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	// Return the product details
	ctx.JSON(http.StatusOK, product)
}
