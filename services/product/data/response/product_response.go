package response

import "v01/model"

type ProductResponse struct {
	Id                int32
	Name              string
	Description       string
	AvailableQuantity float64
	Price             float64
	Category          model.Category
}
