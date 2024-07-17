package response

import "encoding/json"

type PurchaseResponse struct {
	ProductId   int     `json:"productId"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    float64 `json:"quantity"`
}

func (r PurchaseResponse) String() string {
	j, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(j)
}
