package main

import (
	"github.com/alexsandroocanha/Golang-Api-Crud/controller"
	"github.com/alexsandroocanha/Golang-Api-Crud/db"
	"github.com/alexsandroocanha/Golang-Api-Crud/repository"
	"github.com/alexsandroocanha/Golang-Api-Crud/usecase"
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	productRepository := repository.NewProductRepository(dbConnection)

	ProductUseCase := usecase.NewProductUsecase(productRepository)

	productController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", productController.GetProducts)
	server.GET("/product/:productId", productController.GetProductById)
	server.POST("/product", productController.CreateProduct)

	server.Run(":8000")

}
