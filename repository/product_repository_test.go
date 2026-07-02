package repository

import (
	"database/sql"
	"errors"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/alexsandroocanha/Golang-Api-Crud/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetProducts_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "product_name", "price"}).
		AddRow(1, "Mouse", 99.90)
	mock.ExpectQuery("SELECT id, product_name, price FROM product").WillReturnRows(rows)

	repo := NewProductRepository(db)
	products, err := repo.GetProducts()

	require.NoError(t, err)
	assert.Len(t, products, 1)
}

func TestGetProducts_EmptyResult(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "product_name", "price"}) // sem AddRow
	mock.ExpectQuery("SELECT id, product_name, price FROM product").WillReturnRows(rows)

	repo := NewProductRepository(db)
	products, err := repo.GetProducts()

	require.NoError(t, err)
	assert.Empty(t, products)
}

func TestGetProducts_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT id, product_name, price FROM product").
		WillReturnError(errors.New("connection refused"))

	repo := NewProductRepository(db)
	_, err = repo.GetProducts()

	assert.Error(t, err)
}

func TestCreateProduct_PrepareError(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare("INSERT INTO product").WillReturnError(errors.New("syntax error"))

	repo := NewProductRepository(db)
	_, err = repo.CreateProduct(model.Product{Name: "Mouse", Price: 10})

	assert.Error(t, err)
}

func TestGetProductById_NoRows(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare("SELECT \\* FROM product WHERE id = \\$1").
		ExpectQuery().
		WithArgs(999).
		WillReturnError(sql.ErrNoRows)

	repo := NewProductRepository(db)
	product, err := repo.GetProductById(999)

	// isso documenta o bug: err vem preenchido, product NUNCA é nil
	assert.Error(t, err)
	assert.NotNil(t, product)
}
