package main

import (
	"github.com/roblaszczak/go-cleanarch/tests/valid-simple/app"
	"github.com/roblaszczak/go-cleanarch/tests/valid-simple/domain"
	"github.com/roblaszczak/go-cleanarch/tests/valid-simple/infrastructure"
)

func main() {
	payment := app.Payment{Amount: 42.16, Order: domain.Order{}}
	repo := infrastructure.MysqlPaymentsRepository{}

	err := repo.AddPayment(payment)
	if err != nil {
		panic(err)
	}
}
