package request

type ProductRequest struct {
	Id                int32   `json:"id"`
	Name              string  `json:"name" validate:"required"`
	Description       string  `json:"description" validate:"required"`
	AvailableQuantity float64 `json:"availableQuantity" validate:"required,gt=0"`
	Price             float64 `json:"price" validate:"required,gt=0"`
	CategoryId        int32   `json:"categoryId" validate:"required"`
}
