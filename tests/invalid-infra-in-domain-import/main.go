package main

import (
	"github.com/roblaszczak/go-cleanarch/tests/invalid-infra-in-domain-import/app"
	"github.com/roblaszczak/go-cleanarch/tests/invalid-infra-in-domain-import/infrastructure"
)

func main() {
	payment := app.Payment{Amount: 42.16, Order: app.Order{}}
	repo := infrastructure.MysqlPaymentsRepository{}

	err := repo.AddPayment(payment)
	if err != nil {
		panic(err)
	}
}

