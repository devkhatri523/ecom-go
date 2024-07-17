package request

import "encoding/json"

type PurchaseRequest struct {
	ProductId int     `json:"productId"`
	Quantity  float64 `json:"quantity"`
}

func (r PurchaseRequest) String() string {
	j, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(j)
}
