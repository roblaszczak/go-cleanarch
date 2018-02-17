package app_test

import (
	"testing"

	"github.com/roblaszczak/go-cleanarch/examples/invalid-infrastructure-to-app-import-in-tests/app"
	"github.com/roblaszczak/go-cleanarch/examples/invalid-infrastructure-to-app-import-in-tests/infrastructure"
)

func TestNewProduct(t *testing.T) {
	repo := infrastructure.ProductMemoryRepo{}

	product := app.NewProduct()
	repo.Add(product)
}
