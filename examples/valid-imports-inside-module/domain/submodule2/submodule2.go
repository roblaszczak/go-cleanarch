package submodule2

import "github.com/roblaszczak/go-cleanarch/examples/valid-imports-inside-module/domain/submodule1"

type Bar struct {
	foo submodule1.Foo
}
