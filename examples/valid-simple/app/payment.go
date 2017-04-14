package app

import "github.com/roblaszczak/go-cleanarch/examples/valid-simple/domain"

type Payment struct {
	Amount domain.Price
	Order  domain.Order
}

type PaymentsRepository interface {
	AddPayment(Payment) error
}
