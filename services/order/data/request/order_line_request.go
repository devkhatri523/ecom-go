package request

import "encoding/json"

type OrderLineRequest struct {
	Id        int     `json:"id"`
	OrderId   int     `json:"orderId"`
	ProductId int     `json:"productId"`
	Quantity  float64 `json:"quantity"`
}

func (r OrderLineRequest) String() string {
	j, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(j)
}
