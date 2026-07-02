package usecase

import (
	"errors"
	"testing"

	"github.com/alexsandroocanha/Golang-Api-Crud/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct{ mock.Mock }

func (m *MockProductRepository) GetProducts() ([]model.Product, error) {
	args := m.Called()
	return args.Get(0).([]model.Product), args.Error(1)
}

func (m *MockProductRepository) CreateProduct(p model.Product) (int, error) {
	args := m.Called(p)
	return args.Int(0), args.Error(1)
}

func (m *MockProductRepository) GetProductById(id int) (*model.Product, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Product), args.Error(1)
}

func TestCreateProduct_Success(t *testing.T) {
	repo := new(MockProductRepository)
	repo.On("CreateProduct", mock.Anything).Return(10, nil)

	uc := NewProductUsecase(repo)
	result, err := uc.CreateProduct(model.Product{Name: "Cabo HDMI", Price: 29.90})

	assert.NoError(t, err)
	assert.Equal(t, 10, result.ID)
	repo.AssertExpectations(t)
}

func TestCreateProduct_RepositoryError(t *testing.T) {
	repo := new(MockProductRepository)
	repo.On("CreateProduct", mock.Anything).Return(0, errors.New("db error"))

	uc := NewProductUsecase(repo)
	_, err := uc.CreateProduct(model.Product{Name: "Cabo HDMI", Price: 29.90})

	assert.Error(t, err)
}

func TestGetProductById_TableDriven(t *testing.T) {
	tests := []struct {
		name       string
		id         int
		mockReturn *model.Product
		mockErr    error
		expectErr  bool
	}{
		{"produto existe", 1, &model.Product{ID: 1, Name: "Mouse"}, nil, false},
		{"erro no repositorio", 2, nil, errors.New("db down"), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := new(MockProductRepository)
			repo.On("GetProductById", tt.id).Return(tt.mockReturn, tt.mockErr)

			uc := NewProductUsecase(repo)
			_, err := uc.GetProductById(tt.id)

			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
func TestGetProducts_Success(t *testing.T) {
	repo := new(MockProductRepository)
	repo.On("GetProducts").Return([]model.Product{{ID: 1, Name: "Mouse"}}, nil)

	uc := NewProductUsecase(repo)
	products, err := uc.GetProducts()

	assert.NoError(t, err)
	assert.Len(t, products, 1)
	repo.AssertExpectations(t)
}

func TestGetProducts_RepositoryError(t *testing.T) {
	repo := new(MockProductRepository)
	repo.On("GetProducts").Return([]model.Product{}, errors.New("db error"))

	uc := NewProductUsecase(repo)
	_, err := uc.GetProducts()

	assert.Error(t, err)
}
