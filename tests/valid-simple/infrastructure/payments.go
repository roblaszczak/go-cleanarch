package infrastructure

import (
	"fmt"
	"github.com/roblaszczak/go-cleanarch/tests/valid-simple/app"
)

type MysqlPaymentsRepository struct {}

func(r MysqlPaymentsRepository) AddPayment(payment app.Payment) error {
	fmt.Printf("adding payment %+v\n", payment)

	return nil
}