package submodule2

import "github.com/roblaszczak/go-cleanarch/examples/valid-imports-inside-module/domain/submodule1"

// Bar uses submodule1.Foo from domain layer, so it is compatible with Dependency Rule.
type Bar struct {
	foo submodule1.Foo
}
