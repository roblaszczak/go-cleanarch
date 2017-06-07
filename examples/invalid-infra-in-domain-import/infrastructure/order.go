package infrastructure

import (
	"fmt"
)

// MysqlOrderRepository is directly used by domain.OrderAdder.
// Interface should be used here (please take a look for example in OrderAdder docs.
type MysqlOrderRepository struct{}

// AddOrder should be wrapped by an interface.
func (r MysqlOrderRepository) AddOrder(id int) error {
	fmt.Printf("adding order %d", id)

	return nil
}
