package app

import "github.com/roblaszczak/go-cleanarch/examples/valid-simple/domain"

// Payment use some types from domain layer.
type Payment struct {
	Amount domain.Price
	Order  domain.Order
}

// PaymentsRepository is interface, with is implemented in infrastructure layer by MysqlPaymentsRepository.
type PaymentsRepository interface {
	AddPayment(Payment) error
}
