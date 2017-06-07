package domain

import "github.com/roblaszczak/go-cleanarch/examples/ignored-dirs/app"

// Bar is not validated, because all dirs starting with "." are ignored.
type Bar struct {
	foo app.Foo
}
