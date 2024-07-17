package service

import (
	"v01/data/request"
	"v01/data/response"
	"v01/model"
)

type ProductService interface {
	CreateProduct(request request.ProductRequest) (model.Product, error)
	FindAll() ([]response.ProductResponse, error)
	FindById(productId int32) (response.ProductResponse, error)
	Delete(productId int32) (int32, error)
	Update(product request.UpdateProductRequest) error
	PurchaseProducts(purchaseProducts []request.ProductPurchaseRequest) ([]response.ProductPurchaseResponse, error)
}
