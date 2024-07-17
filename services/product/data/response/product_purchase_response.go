package response

type ProductPurchaseResponse struct {
	ProductId   int32   `json:"productId"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    float64 `json:"quantity"`
}
