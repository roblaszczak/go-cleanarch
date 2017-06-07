package submodule

import "github.com/roblaszczak/go-cleanarch/examples/invalid-imports-between-submodules/application/submodule"

// Bar imports submodule.Foo. Only interfaces layer can be used in another modules.
type Bar struct {
	foo submodule.Foo
}
