package enum

type PaymentMethod int

const (
	UNKNOWN = iota
	PAYPAL
	CREDIT_CARD
	VISA
	MASTER_CARD
	BITCOIN
)

func (p PaymentMethod) String() string {
	return [...]string{"UNKNOWN", "PAYPAL", "CREDIT", "VISA", "MASTERCARD", "BITCOIN"}[p]
}
