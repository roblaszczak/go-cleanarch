package interfaces

import "github.com/roblaszczak/go-cleanarch/examples/valid-cross-module-deps/cms/auth/usecases"

func CheckAccess(username string) bool {
	return usecases.LoginAccessChecker(username)
}
