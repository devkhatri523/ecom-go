package repository

import "v01/model"

type ProductRepository interface {
	CreateProduct(product model.Product) (int32, error)
	FindAll() ([]model.Product, error)
	FindById(productId int32) (model.Product, error)
	Delete(productId int32) (int32, error)
	Save(product model.Product) (int32, error)
	Update(product model.Product) error
	FindAllProductsById(ids []int32) ([]model.Product, error)
}
