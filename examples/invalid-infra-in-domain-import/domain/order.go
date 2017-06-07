package domain

import "github.com/roblaszczak/go-cleanarch/examples/invalid-infra-in-domain-import/infrastructure"

// OrderAdder uses infrastructure.MysqlOrderRepository directly, with breaks Dependency Rule.
//
// You should use interface here. Interface should look like this:
//   type OrderRepository interface {
//     AddOrder(id int) error
//   }
type OrderAdder struct {
	infrastructure.MysqlOrderRepository
}

// AddOrder should call repository using interface, not type from infrastructure.
func (a OrderAdder) AddOrder(order Order) error {
	return a.MysqlOrderRepository.AddOrder(order.ID)
}

// Order is used by infrastructure.MysqlOrderRepository, with is good because we don't break Dependency Rule.
type Order struct {
	ID int
}
