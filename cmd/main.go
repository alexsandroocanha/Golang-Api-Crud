package main

import (
	"log"
	"net/http"

	"github.com/alexsandroocanha/Golang-Api-Crud/controller"
	"github.com/alexsandroocanha/Golang-Api-Crud/db"
	"github.com/alexsandroocanha/Golang-Api-Crud/repository"
	"github.com/alexsandroocanha/Golang-Api-Crud/usecase"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
		ctx.JSON(200, gin.H{"message": "pong"})
	})
	server.GET("/products", productController.GetProducts)
	server.GET("/product/:productId", productController.GetProductById)
	server.POST("/product", productController.CreateProduct)

	go func() {
		metricsMux := http.NewServeMux()
		metricsMux.Handle("/metrics", promhttp.Handler())

		log.Println("servidor de métricas rodando em :9100")
		if err := http.ListenAndServe(":9100", metricsMux); err != nil {
			log.Fatalf("erro ao subir servidor de métricas: %v", err)
		}
	}()

	if err := server.Run(":8000"); err != nil {
		log.Fatalf("erro ao subir servidor principal: %v", err)
	}
}
