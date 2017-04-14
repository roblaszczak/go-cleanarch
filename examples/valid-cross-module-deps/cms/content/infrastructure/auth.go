package infrastructure

import (
	"github.com/roblaszczak/go-cleanarch/examples/valid-cross-module-deps/cms/auth/interfaces"
)

type AuthModuleAuthChecker struct{}

func (a AuthModuleAuthChecker) CheckAuth(username string) bool {
	return interfaces.CheckAccess(username)
}
