package domain

type PaymentNotification struct {
	OrderReference    string  `json:"orderReference"`
	Amount            float64 `json:"amount"`
	CustomerFirstName string  `json:"customerFirstName"`
	CustomerLastName  string  `json:"customerLastName"`
	PaymentMethod     string  `json:"paymentMethod"`
	CustomerEmail     string  `json:"customerEmail"`
}
