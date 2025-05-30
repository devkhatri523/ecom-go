package domain

import (
	"time"
)

type Order struct {
	Id               int32     `gorm:"column:id;primaryKey"`
	Reference        string    `gorm:"column:reference"`
	TotalAmount      float64   `gorm:"column:total_amount"`
	PaymentMethod    string    `gorm:"column:payment_method"`
	CustomerId       string    `gorm:"column:customer_id"`
	CreatedAt        time.Time `gorm:"column:created_at"`
	LastModifiedDate time.Time `gorm:"column:last_modified_date"`
}

func (p *Order) TableName() string {
	return "customer_order"
}
