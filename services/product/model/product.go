package model

type Product struct {
	Id                 int32    `gorm:"primaryKey"`
	ProductName        string   `gorm:"column:name"`
	ProductDescription string   `gorm:"column:description"`
	AvailableQuantity  float64  `gorm:"column:available_quantity"`
	Price              float64  `gorm:"column:price"`
	CategoryId         int32    `gorm:"column:category_id"`
	Category           Category `gorm:"foreignKey:CategoryId;references:CategoryId"`
}

func (p *Product) TableName() string {
	return "product"
}

type ProductCategory struct {
	Id                  int32
	ProductName         string
	ProductDescription  string
	AvailableQuantity   float64
	Price               float64
	CategoryID          int32
	CategoryName        string
	CategoryDescription string
}
