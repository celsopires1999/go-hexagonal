package application_test

import (
	"testing"

	"github.com/celsopires1999/go-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	// product := application.Product{}
	product := application.NewProduct()
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	// product := application.Product{}
	product := application.NewProduct()
	product.Name = "Hello"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be zero in order to have the product disabled", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	// product := application.Product{}
	product := application.NewProduct()
	product.ID = uuid.NewV4().String()
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater than or equal zero", err.Error())
}

func TestProduct_GetID(t *testing.T) {
	uuidTest := uuid.NewV4().String()
	product := application.Product{}
	product.ID = uuidTest

	require.Equal(t, uuidTest, product.GetID())
}

func TestProduct_GetName(t *testing.T) {
	nameTest := "Car"
	// product := application.Product{}
	product := application.NewProduct()
	product.Name = nameTest

	require.Equal(t, nameTest, product.GetName())
}

func TestProduct_GetStatus(t *testing.T) {
	// product := application.Product{}
	product := application.NewProduct()
	product.Status = application.ENABLED

	require.Equal(t, application.ENABLED, product.GetStatus())
}

func TestProduct_GetPrice(t *testing.T) {
	// product := application.Product{}
	product := application.NewProduct()
	product.Price = 10.0

	require.Equal(t, 10.0, product.GetPrice())
}
