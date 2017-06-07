package infrastructure

import (
	"github.com/roblaszczak/go-cleanarch/examples/valid-cross-module-deps/cms/auth/interfaces"
)

// AuthModuleAuthChecker has method CheckAuth, with is using function from another module interface.
type AuthModuleAuthChecker struct{}

// CheckAuth is using interfaces.CheckAccess from another module, with is valida accoding to Dependency Rule.
func (a AuthModuleAuthChecker) CheckAuth(username string) bool {
	return interfaces.CheckAccess(username)
}
