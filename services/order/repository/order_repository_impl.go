package repository

import (
	"errors"
	"fmt"
	"github.com/devkhatri523/order-service/domain"
	"github.com/devkhatri523/order-service/helper"
	"github.com/devkhatri523/order-service/query"
	"gorm.io/gorm"
	"log"
)

type OrderRepositoryImpl struct {
	Db *gorm.DB
}

func NewOrderRepositoryImpl(Dd *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{
		Db: Dd,
	}
}

func (o OrderRepositoryImpl) CreateOrder(order domain.Order) (int32, error) {
	result := o.Db.Create(&order)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return order.Id, nil
}

func (o OrderRepositoryImpl) FindAllOrder(filter *domain.Filters) ([]*domain.Order, *domain.Pagination, error) {
	p := helper.CreateGetAllOrderPagination(filter.Pagination)
	pQuery, err := p.GetPaginationQuery(o.Db, []*domain.Order{})
	if err != nil {
		return nil, nil, err
	}
	stmt := o.Db.Raw(pQuery)
	var orders []*domain.Order
	results := stmt.Find(&orders)
	pagination := new(domain.Pagination)
	if results.Error != nil {
		return nil, pagination, results.Error
	}
	cursor, err := p.GetCursor(&orders)
	if err != nil {
		return nil, nil, err
	}
	pagination = &domain.Pagination{}
	if cursor.After != nil {
		pagination.AfterCursor = *cursor.After
	}
	if cursor.Before != nil {
		pagination.BeforeCursor = *cursor.Before
	}
	return orders, pagination, nil
}

func (o OrderRepositoryImpl) FindByOrderId(orderId int32) (domain.Order, error) {
	var order domain.Order
	whereQuery := fmt.Sprintf("%s where a.id = %d", query.ORDER_QUERY, orderId)
	stmt := o.Db.Raw(whereQuery)
	results := stmt.Find(&order)
	if results.Error != nil {
		return domain.Order{}, results.Error
	}
	if results.RowsAffected == 0 {
		return domain.Order{}, errors.New("Order is not found")
	}
	return order, nil
}
