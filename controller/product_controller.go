package controller

import (
	"net/http"

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
	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}
