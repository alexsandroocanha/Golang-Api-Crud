package main

import (
	"github.com/alexsandroocanha/Golang-Api-Crud/controller"
	"github.com/alexsandroocanha/Golang-Api-Crud/usecase"
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	ProductUseCase := usecase.NewProductUsecase()

	productController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", productController.GetProducts)

	server.Run(":8000")

}
