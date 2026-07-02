package usecase

import (
	"github.com/alexsandroocanha/Golang-Api-Crud/model"
	"github.com/alexsandroocanha/Golang-Api-Crud/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepositoryInterface
}

func NewProductUsecase(repo repository.ProductRepositoryInterface) *ProductUsecase {
	return &ProductUsecase{
		repository: repo,
	}
}

type ProductUsecaseInterface interface {
	GetProducts() ([]model.Product, error)
	CreateProduct(product model.Product) (model.Product, error)
	GetProductById(id_product int) (*model.Product, error)
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {

	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil

}
func (pu *ProductUsecase) GetProductById(id_product int) (*model.Product, error) {

	product, err := pu.repository.GetProductById(id_product)
	if err != nil {
		return nil, err
	}

	return product, nil
}
