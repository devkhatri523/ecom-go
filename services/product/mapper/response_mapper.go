package mapper

import (
	"v01/data/response"
	"v01/model"
)

/*func MapToProductResponse(productList []model.Product) []response.ProductResponse {
	var products []response.ProductResponse

	for _, value := range productList {
		product := response.ProductResponse{
			Name:              value.Name,
			Description:       value.Description,
			Id:                value.Id,
			AvailableQuantity: value.AvailableQuantity,
		}

	}
}*/

func MapToPurchaseResponse(product model.Product, quantity float64) response.ProductPurchaseResponse {
	productPurchaseReponse := response.ProductPurchaseResponse{
		ProductId:   product.Id,
		Name:        product.ProductName,
		Description: product.ProductDescription,
		Price:       product.Price,
		Quantity:    quantity,
	}
	return productPurchaseReponse

}
