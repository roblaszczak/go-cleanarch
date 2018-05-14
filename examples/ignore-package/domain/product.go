package domain

import "github.com/roblaszczak/go-cleanarch/examples/ignore-package/app"

// Product imports app.Price, with breaks Dependency Rule.
type Product struct {
	Price app.Price
}
