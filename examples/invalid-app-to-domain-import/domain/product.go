package domain

import "github.com/roblaszczak/go-cleanarch/examples/invalid-app-to-domain-import/app"

// Product imports app.Price, with breaks Dependency Rule.
type Product struct {
	Price app.Price
}
