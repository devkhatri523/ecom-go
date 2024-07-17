package service

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"sort"
	"v01/data/request"
	"v01/data/response"
	"v01/mapper"
	"v01/model"
	"v01/repository"
)

type ProductServiceImpl struct {
	repository.ProductRepository
	Validate *validator.Validate
}

func NewProductServiceImpl(respository repository.ProductRepository, validate *validator.Validate) (service ProductService, err error) {
	if validate == nil {
		return nil, errors.New("validator instance cannot be nil")
	}

	return &ProductServiceImpl{
		ProductRepository: respository,
		Validate:          validate,
	}, err
}

func (p ProductServiceImpl) CreateProduct(request request.ProductRequest) (model.Product, error) {
	err := p.Validate.Struct(request)
	if err != nil {
		return model.Product{}, err
	}
	m := model.Product{
		Id:                 request.Id,
		ProductName:        request.Name,
		ProductDescription: request.Description,
		AvailableQuantity:  request.AvailableQuantity,
		Price:              request.Price,
		CategoryId:         request.CategoryId,
		Category: model.Category{
			CategoryId: request.CategoryId,
		},
	}
	_, err = p.ProductRepository.Save(m)
	if err != nil {
		return model.Product{}, err
	}
	return m, nil
}
func (p ProductServiceImpl) FindAll() (productResponse []response.ProductResponse, err error) {
	result, err := p.ProductRepository.FindAll()
	if err != nil {
		return nil, err
	}
	fmt.Println(result)

	for _, value := range result {
		product := response.ProductResponse{
			Id:          value.Id,
			Name:        value.ProductName,
			Description: value.ProductDescription,
			Category:    value.Category,
		}
		productResponse = append(productResponse, product)
	}
	return productResponse, nil

}

func (p ProductServiceImpl) FindById(productId int32) (response.ProductResponse, error) {
	data, err := p.ProductRepository.FindById(productId)
	if err != nil {
		return response.ProductResponse{}, err
	}

	res := response.ProductResponse{
		Id:          data.Id,
		Name:        data.ProductName,
		Description: data.ProductDescription,
		Category:    data.Category,
	}
	return res, nil
}

func (p ProductServiceImpl) Delete(productId int32) (int32, error) {
	productId, err := p.ProductRepository.Delete(productId)

	if err != nil {
		return 0, err
	}

	return productId, nil
}

func (p ProductServiceImpl) Update(req request.UpdateProductRequest) error {
	value := model.Product{
		Id:                 req.Id,
		ProductName:        req.Name,
		ProductDescription: req.Description,
		AvailableQuantity:  float64(req.AvailableQuantity),
	}
	err := p.ProductRepository.Update(value)
	if err != nil {
		return err
	}
	return nil
}

func (p ProductServiceImpl) PurchaseProducts(purchaseProductRequest []request.ProductPurchaseRequest) ([]response.ProductPurchaseResponse, error) {
	var productIds []int32
	var purchaseProductResponse []response.ProductPurchaseResponse
	for _, productRequest := range purchaseProductRequest {
		productIds = append(productIds, productRequest.Id)
	}
	storedProducts, err := p.ProductRepository.FindAllProductsById(productIds)
	if err != nil {
		return nil, errors.New("exception occurred while purchasing products")
	}
	if len(productIds) != len(storedProducts) {
		return nil, errors.New("one or more product does not exists")
	}
	sort.Slice(purchaseProductRequest, func(i, j int) bool {
		return purchaseProductRequest[i].Id > purchaseProductRequest[j].Id
	})
	for i := 0; i < len(storedProducts); i++ {
		product := storedProducts[i]
		productRequest := purchaseProductRequest[i]
		if product.AvailableQuantity < productRequest.Quantity {

			return nil, errors.New(fmt.Sprintf("Insufficient stock quantity for product with Id: %d", product.Id))
		}
		newAvailableQuantity := product.AvailableQuantity - productRequest.Quantity
		product.AvailableQuantity = newAvailableQuantity
		_, err = p.ProductRepository.Save(product)
		if err != nil {
			return nil, err
		}
		purchaseProductResponse = append(purchaseProductResponse, mapper.MapToPurchaseResponse(product, productRequest.Quantity))
	}
	return purchaseProductResponse, nil

}
