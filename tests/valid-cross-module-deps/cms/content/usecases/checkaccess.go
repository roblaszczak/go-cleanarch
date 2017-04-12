package usecases

import (
	"github.com/roblaszczak/go-cleanarch/tests/valid-cross-module-deps/cms/content/domain"
)

type AccessCheckController struct {
	AuthChecker domain.AuthChecker
}

func (c AccessCheckController) CheckAccess(username string) bool {
	return c.AuthChecker.CheckAuth(domain.NewUser(username))
}
