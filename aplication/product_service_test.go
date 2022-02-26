package aplication_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go-hexa/aplication"
	mock_aplication "go-hexa/aplication/mocks"
	"testing"
)

func TestProductService_Get(t *testing.T) {
	controller := gomock.NewController(t)

	defer controller.Finish()
	product := mock_aplication.NewMockProductInterface(controller)
	productRepository := mock_aplication.NewMockProductRepository(controller)

	productRepository.EXPECT().FindById(gomock.Any()).Return(product, nil).AnyTimes()

	service := aplication.ProductService{
		productRepository,
	}

	result, err := service.Get("id")
	require.Nil(t, err)
	require.Equal(t, result, product)

}
func TestProductService_Create(t *testing.T) {
	controller := gomock.NewController(t)

	defer controller.Finish()
	product := mock_aplication.NewMockProductInterface(controller)
	productRepository := mock_aplication.NewMockProductRepository(controller)

	productRepository.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := aplication.ProductService{
		ProductRepository: productRepository,
	}

	result, err := service.Create("name", 10)
	require.Nil(t, err)
	require.Equal(t, result, product)

}

func TestProductService_Enable(t *testing.T) {
	controller := gomock.NewController(t)

	defer controller.Finish()
	product := mock_aplication.NewMockProductInterface(controller)
	productRepository := mock_aplication.NewMockProductRepository(controller)
	product.EXPECT().Enable().Return(nil)

	productRepository.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := aplication.ProductService{
		ProductRepository: productRepository,
	}

	result, err := service.Enable(product)

	require.Nil(t, err)
	require.Equal(t, result, product)

}
func TestProductService_Disable(t *testing.T) {
	controller := gomock.NewController(t)

	defer controller.Finish()
	product := mock_aplication.NewMockProductInterface(controller)
	productRepository := mock_aplication.NewMockProductRepository(controller)
	product.EXPECT().Disable().Return(nil)

	productRepository.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := aplication.ProductService{
		ProductRepository: productRepository,
	}

	result, err := service.Disable(product)

	require.Nil(t, err)
	require.Equal(t, result, product)

}
