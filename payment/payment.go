package payment

type Payment int

type Type struct {
	PaymentType Payment `json:"payment_type"`
	Price       float64 `json:"price"`
}

const (
	Free Payment = iota
	Paid
	Subscription
)
