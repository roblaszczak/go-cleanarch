package infrastructure

import "github.com/roblaszczak/go-cleanarch/examples/invalid-infrastructure-to-app-import-in-tests/app"

// ProductMemoryRepo is imported in domain.Product tests.
type ProductMemoryRepo struct {}

func (p *ProductMemoryRepo) Add(product *app.Product) {
	// ...
}
