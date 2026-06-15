package controller

import (
	"net/http"

	"github.com/alexsandroocanha/Golang-Api-Crud/model"
	"github.com/gin-gonic/gin"
)

type productController struct {
}

func NewProductController() productController {
	return productController{}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		ID:    1,
		Name:  "Batata frita",
		Price: 300,
	}

	ctx.JSON(http.StatusOK, products)
}
