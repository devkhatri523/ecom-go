package repository

import (
	"errors"
	"fmt"
	"github.com/devkhatri523/ecom-go/go-service/logger"
	"github.com/devkhatri523/order-service/domain"
	"github.com/devkhatri523/order-service/query"
	"gorm.io/gorm"
)

type OrderLineRepositoryImpl struct {
	Db *gorm.DB
}

func NewOrderLineRepositoryImpl(Db *gorm.DB) OrderLineRepository {
	return &OrderLineRepositoryImpl{Db: Db}
}

func (o OrderLineRepositoryImpl) SaveOrderLine(orderLine domain.OrderLine) (int32, error) {
	result := o.Db.Create(&orderLine)
	if result.Error != nil {
		logger.Default().Error(result.Error)
	}
	return orderLine.Id, nil
}

func (o OrderLineRepositoryImpl) FindAllByOrderId(orderLineId int32) (domain.OrderLine, error) {
	var orderLine domain.OrderLine
	whereQuery := fmt.Sprintf("%s where a.id = %d", query.ORDER_QUERY, orderLineId)
	stmt := o.Db.Raw(whereQuery)
	results := stmt.Find(&orderLine)
	if results.Error != nil {
		return domain.OrderLine{}, results.Error
	}
	if results.RowsAffected == 0 {
		return domain.OrderLine{}, errors.New("Order is not found")
	}
	return orderLine, nil
}
