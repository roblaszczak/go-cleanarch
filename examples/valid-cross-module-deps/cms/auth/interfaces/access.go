package interfaces

import "github.com/roblaszczak/go-cleanarch/examples/valid-cross-module-deps/cms/auth/usecases"

// CheckAccess use usecases.LoginAccessChecker, with is compatible with Dependency Rule.
//
// CheckAccess is also used by usecases.AddArticle in another module,
// with is also valid because we can use interfaces layer of module in another module.
func CheckAccess(username string) bool {
	return usecases.LoginAccessChecker(username)
}
