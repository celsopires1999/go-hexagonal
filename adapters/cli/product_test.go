package cli_test

import (
	"fmt"
	"testing"

	"github.com/celsopires1999/go-hexagonal/adapters/cli"
	mock_application "github.com/celsopires1999/go-hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	crtl := gomock.NewController(t)
	defer crtl.Finish()

	productName := "Product Test"
	productPrice := 25.99
	productStatus := "enabled"
	productId := "abc"

	// mock of Product
	productMock := mock_application.NewMockProductInterface(crtl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()

	// mock of Service
	service := mock_application.NewMockProductServiceInterface(crtl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	// Test of create
	resultExpected := fmt.Sprintf("Product ID # %s with the name %s has been created with the price %f and status %s",
		productId, productName, productPrice, productStatus)
	result, err := cli.Run(service, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	// Test of enable
	resultExpected = fmt.Sprintf("Product %s has been enabled.", productName)
	result, err = cli.Run(service, "enable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	// Test of disable
	resultExpected = fmt.Sprintf("Product %s has been disabled.", productName)
	result, err = cli.Run(service, "disable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	// Test of get
	resultExpected = fmt.Sprintf("Product ID: %s\nName: %s\nPrice:%f\nStatus: %s",
		productId, productName, productPrice, productStatus)
	result, err = cli.Run(service, "get", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}
