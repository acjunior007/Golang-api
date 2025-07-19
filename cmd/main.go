package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// Initialize database connection
	db, err := db.ConnectDB()
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	// Initialize repositories
	productRepo := repository.NewProductRepository(db)

	// Initialize usecases
	productUseCase := usecase.NewProductUsecase(productRepo)

	// Initialize controllers
	productController := controller.NewProductController(productUseCase)

	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Ok!",
		})
	})

	server.GET("/products", productController.GetProducts)

	server.POST("/products", productController.CreateProduct)

	server.GET("/products/:id", productController.GetProductByID)

	server.DELETE("/products/:id", productController.Delete)

	server.Run(":8000")
}
