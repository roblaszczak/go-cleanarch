package domain

import "github.com/roblaszczak/go-cleanarch/examples/invalid-infra-in-domain-import/infrastructure"

type OrderAdder struct {
	// please, don't do this :(
	// you should use an interface here
	infrastructure.MysqlOrderRepository
}

func (a OrderAdder) AddOrder(order Order) error {
	return a.MysqlOrderRepository.AddOrder(order.Id)
}

type Order struct {
	Id int
}
