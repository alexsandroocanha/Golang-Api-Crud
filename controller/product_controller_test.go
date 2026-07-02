package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/alexsandroocanha/Golang-Api-Crud/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProductUsecase struct{ mock.Mock }

func (m *MockProductUsecase) GetProducts() ([]model.Product, error) {
	args := m.Called()
	return args.Get(0).([]model.Product), args.Error(1)
}

func (m *MockProductUsecase) CreateProduct(p model.Product) (model.Product, error) {
	args := m.Called(p)
	return args.Get(0).(model.Product), args.Error(1)
}

func (m *MockProductUsecase) GetProductById(id int) (*model.Product, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Product), args.Error(1)
}

func TestGetProducts_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	uc := new(MockProductUsecase)
	uc.On("GetProducts").Return([]model.Product{{ID: 1, Name: "Mouse"}}, nil)

	ctrl := NewProductController(uc)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/products", nil)

	ctrl.GetProducts(c)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetProducts_UsecaseError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	uc := new(MockProductUsecase)
	uc.On("GetProducts").Return([]model.Product{}, errors.New("falha"))

	ctrl := NewProductController(uc)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/products", nil)

	ctrl.GetProducts(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestCreateProduct_InvalidBody(t *testing.T) {
	gin.SetMode(gin.TestMode)
	uc := new(MockProductUsecase)
	ctrl := NewProductController(uc)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/product", strings.NewReader(`{invalid`))
	c.Request.Header.Set("Content-Type", "application/json")

	ctrl.CreateProduct(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateProduct_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	uc := new(MockProductUsecase)
	uc.On("CreateProduct", mock.Anything).Return(model.Product{ID: 3, Name: "Mouse", Price: 50}, nil)

	ctrl := NewProductController(uc)
	body := `{"name":"Mouse","price":50}`
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/product", strings.NewReader(body))
	c.Request.Header.Set("Content-Type", "application/json")

	ctrl.CreateProduct(c)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetProductById_NonNumericId(t *testing.T) {
	gin.SetMode(gin.TestMode)
	uc := new(MockProductUsecase)
	ctrl := NewProductController(uc)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/product/abc", nil)
	c.Params = gin.Params{{Key: "productId", Value: "abc"}}

	ctrl.GetProductById(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetProductById_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	uc := new(MockProductUsecase)
	uc.On("GetProductById", 1).Return(&model.Product{ID: 1, Name: "Mouse"}, nil)

	ctrl := NewProductController(uc)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodGet, "/product/1", nil)
	c.Params = gin.Params{{Key: "productId", Value: "1"}}

	ctrl.GetProductById(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var product model.Product
	json.Unmarshal(w.Body.Bytes(), &product)
	assert.Equal(t, "Mouse", product.Name)
}
