package domain

type OrderLine struct {
	Id        int32   `gorm:"column:id;primaryKey"`
	ProductId int32   `gorm:"column:product_id"`
	Quantity  float64 `gorm:"column:quantity"`
	OrderId   int32   `gorm:"column:order_id"`
}

func (p *OrderLine) TableName() string {
	return "order_line"
}
