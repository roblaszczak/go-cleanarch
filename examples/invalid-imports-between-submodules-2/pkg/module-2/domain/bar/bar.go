package bar

import "github.com/roblaszczak/go-cleanarch/examples/invalid-imports-between-submodules-2/pkg/module-1/interfaces/foo"

// Bar imports foo.Foo. We can only include interfaces to infrastructure layer.
type Bar struct {
	foo foo.Foo
}
