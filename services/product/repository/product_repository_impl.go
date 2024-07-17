package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"v01/model"
	"v01/query"
)

type ProductRepositoryImpl struct {
	Db *gorm.DB
}

func NewProductRepositoryImpl(Db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{
		Db: Db,
	}

}

func (p ProductRepositoryImpl) CreateProduct(product model.Product) (int32, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepositoryImpl) FindAll() (products []model.Product, err error) {
	var productCategory []model.ProductCategory
	stmt := p.Db.Raw(query.PRODUCT_QUERY)
	results := stmt.Find(&productCategory)

	for _, value := range productCategory {
		product := model.Product{
			Id:                 value.Id,
			ProductName:        value.ProductName,
			ProductDescription: value.ProductDescription,
			AvailableQuantity:  value.AvailableQuantity,
			Category: model.Category{
				CategoryName:        value.CategoryName,
				CategoryId:          value.CategoryID,
				CategoryDescription: value.CategoryDescription,
			},
		}
		products = append(products, product)

	}

	if results.Error != nil {
		return nil, results.Error
	}

	return products, nil
}

func (p ProductRepositoryImpl) FindById(productId int32) (model.Product, error) {
	var productCategory model.ProductCategory
	whereQuery := fmt.Sprintf("%s where a.id = %d", query.PRODUCT_QUERY, productId)
	stmt := p.Db.Raw(whereQuery)
	results := stmt.Find(&productCategory)
	if results.Error != nil {
		return model.Product{}, results.Error
	}
	if results.RowsAffected == 0 {
		return model.Product{}, errors.New("Product is not found")
	}

	product := model.Product{
		Id:                 productCategory.Id,
		ProductName:        productCategory.ProductName,
		ProductDescription: productCategory.ProductDescription,
		AvailableQuantity:  productCategory.AvailableQuantity,
		Category: model.Category{
			CategoryName:        productCategory.CategoryName,
			CategoryId:          productCategory.CategoryID,
			CategoryDescription: productCategory.CategoryDescription,
		},
	}

	return product, nil
}

func (p ProductRepositoryImpl) Delete(productId int32) (int32, error) {
	var product model.Product

	result := p.Db.Where("id = ?", productId).Delete(&product)
	if result.Error != nil {
		return 0, result.Error
	}
	return productId, nil
}

func (p ProductRepositoryImpl) Save(product model.Product) (int32, error) {
	result := p.Db.Create(&product)
	if result.Error != nil {
		return 0, result.Error
	}

	return product.Id, nil
}

func (p ProductRepositoryImpl) Update(product model.Product) error {
	data := &model.Product{
		ProductName:        product.ProductName,
		ProductDescription: product.ProductDescription,
		AvailableQuantity:  product.AvailableQuantity,
	}

	result := p.Db.Where(&model.Product{Id: product.Id}).UpdateColumns(&data)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p ProductRepositoryImpl) FindAllProductsById(ids []int32) ([]model.Product, error) {
	var productCategories []model.ProductCategory
	var productList []model.Product
	whereQuery := fmt.Sprintf("%s where a.id IN( %s ) order by id", query.PRODUCT_QUERY, buildInClause(ids))
	stmt := p.Db.Raw(whereQuery)
	results := stmt.Find(&productCategories)
	if results.Error != nil {
		return nil, results.Error
	}
	if results.RowsAffected == 0 {
		return nil, errors.New("Product is not found")
	}

	for _, indvProductCategory := range productCategories {
		product := model.Product{
			Id:                 indvProductCategory.Id,
			ProductName:        indvProductCategory.ProductName,
			ProductDescription: indvProductCategory.ProductDescription,
			AvailableQuantity:  indvProductCategory.AvailableQuantity,
			Category: model.Category{
				CategoryName:        indvProductCategory.CategoryName,
				CategoryId:          indvProductCategory.CategoryID,
				CategoryDescription: indvProductCategory.CategoryDescription,
			},
		}
		productList = append(productList, product)
	}

	return productList, nil
}

func buildInClause(ints []int32) string {
	strs := make([]string, len(ints))
	for i, v := range ints {
		strs[i] = strconv.Itoa(int(v))
	}
	return strings.Join(strs, ", ")
}
