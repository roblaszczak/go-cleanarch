package infrastructure

import (
	"github.com/roblaszczak/go-cleanarch/tests/valid-cross-module-deps/cms/auth/interfaces"
	"github.com/roblaszczak/go-cleanarch/tests/valid-cross-module-deps/cms/content/domain"
)

type AuthModuleAuthChecker struct{}

func (a AuthModuleAuthChecker) CheckAuth(u domain.User) bool {
	return interfaces.CheckAccess(u.Username())
}
