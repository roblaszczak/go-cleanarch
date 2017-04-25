package submodule

import "github.com/roblaszczak/go-cleanarch/examples/invalid-imports-between-submodules/application/submodule"

type Bar struct {
	foo submodule.Foo
}