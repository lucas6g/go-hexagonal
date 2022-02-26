package cli

import (
	"fmt"
	"go-hexa/aplication"
)

func Run(service aplication.ProductServiceInterface, action string, productId string, productName string, productPrice float32) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, productPrice)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product id %s whit the name %s has bean created whit the price %f  and status %s ",
			product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus())

	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		product, err = service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product whit name %s has been enable ",
			product.GetName())

	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		product, err = service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product whit name %s has been disabled ",
			product.GetName())

	default:
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("%s\nId, %s\nName,  %f\nPrice, %s\nStatus",
			product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus())

	}

	return result, nil

}
