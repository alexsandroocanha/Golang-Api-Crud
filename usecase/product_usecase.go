package usecase

import "github.com/alexsandroocanha/Golang-Api-Crud/model"

type ProductUsecase struct {
	// repository
}

func NewProductUsecase() ProductUsecase {
	return ProductUsecase{}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return []model.Product{}, nil
}
