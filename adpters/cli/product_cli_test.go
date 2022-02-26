package cli_test

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go-hexa/adpters/cli"
	mock_aplication "go-hexa/aplication/mocks"
	"testing"
)

func TestRun(t *testing.T) {
	controller := gomock.NewController(t)

	defer controller.Finish()

	productName := "Product test"
	productPrice := float32(25)
	productStatus := "enabled"
	productId := "uuid"

	productMock := mock_aplication.NewMockProductInterface(controller)

	//mockando o produto que vai ser retornado no cli
	productMock.EXPECT().GetId().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	service := mock_aplication.NewMockProductServiceInterface(controller)

	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product id %s whit the name %s has bean created whit the price %f  and status %s ",
		productId, productName, productPrice, productStatus)

	//metodo create

	result, err := cli.Run(service, "create", "", productName, productPrice)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	// metodo disabled
	resultExpected = fmt.Sprintf("Product whit name %s has been disabled ", productName)

	result, err = cli.Run(service, "disable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	// metodo enable
	resultExpected = fmt.Sprintf("Product whit name %s has been enable ", productName)

	result, err = cli.Run(service, "enable", productId, "", 0)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	// returns product when no passes a acction
	resultExpected = fmt.Sprintf("%s\nId, %s\nName,  %f\nPrice, %s\nStatus", productId, productName, productPrice, productStatus)

	result, err = cli.Run(service, "", productId, "", 0)

	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

}
