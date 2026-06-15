package controller

import (
	"net/http"

	"github.com/alexsandroocanha/Golang-Api-Crud/model"
	"github.com/alexsandroocanha/Golang-Api-Crud/usecase"
	"github.com/gin-gonic/gin"
)

type productController struct {
	productUseCase usecase.ProductUsecase
}

func NewProductController(useCase usecase.ProductUsecase) productController {
	return productController{
		productUseCase: useCase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		{
			ID:    1,
			Name:  "Batata frita",
			Price: 300,
		},
		{
			ID:    2,
			Name:  "Milk Shake",
			Price: 250,
		},
	}

	ctx.JSON(http.StatusOK, products)
}
