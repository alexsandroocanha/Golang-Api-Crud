package usecase

import (
	"github.com/alexsandroocanha/Golang-Api-Crud/model"
	"github.com/alexsandroocanha/Golang-Api-Crud/repository"
)

type ProductUsecase struct {
	// repository
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}
