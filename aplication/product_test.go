package aplication_test

import (
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"go-hexa/aplication"
	"testing"
)

func TestProduct_Enabled(t *testing.T) {
	product := aplication.Product{}
	product.Name = "anyName"
	product.Status = "disabled"
	product.Price = 10

	err := product.Enable()

	//nao deve retornar um erro espero um nil
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()

	require.Equal(t, "the price most be big then zero", err.Error())
}

func TestProduct_Disabled(t *testing.T) {
	product := aplication.Product{}
	product.Name = "anyName"
	product.Status = aplication.DISABLED
	product.Price = 0

	err := product.Disable()

	//nao deve retornar um erro espero um nil
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()

	require.Equal(t, "the price most be equal to zero", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := aplication.Product{}
	product.Id = uuid.NewV4().String()
	product.Name = "anyName"
	product.Status = aplication.DISABLED
	product.Price = 10

	_, err := product.IsValid()

	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()

	require.Equal(t, "status most be enebled or desable", err.Error())

	product.Status = ""
	_, err = product.IsValid()

	require.Equal(t, product.Status, aplication.DISABLED)

	product.Price = -10
	_, err = product.IsValid()

	require.Equal(t, "the price most be 0 or big than zero", err.Error())

}
