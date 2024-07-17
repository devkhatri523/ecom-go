package request

type UpdateProductRequest struct {
	Id                int32
	Name              string `validate:"required,max=200,min=1" json:"name"`
	Description       string `validate:"required,max=200,min=1" json:"description"`
	AvailableQuantity int32  `validate:"required,max=200,min=1" json:"availableQuantity"`
}
