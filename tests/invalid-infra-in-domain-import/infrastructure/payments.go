package infrastructure

import (
	"fmt"
)

type MysqlOrderRepository struct {}

func(r MysqlOrderRepository) AddOrder(id int) error {
	fmt.Printf("adding order %d", id)

	return nil
}