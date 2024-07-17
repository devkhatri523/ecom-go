package model

type Category struct {
	CategoryId          int32  `gorm:"primaryKey;column:id"`
	CategoryName        string `gorm:"column:name"`
	CategoryDescription string `gorm:"column:description"`
}

func (p *Category) TableName() string {
	return "category"
}
