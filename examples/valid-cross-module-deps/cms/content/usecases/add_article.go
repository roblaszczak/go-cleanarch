package usecases

import (
	"fmt"
)

// AuthChecker is interface comptible with AuthModuleAuthChecker.
type AuthChecker interface {
	CheckAuth(username string) bool
}

// AddArticle can add article, and check authentication using  AuthChecker.
type AddArticle struct {
	AuthChecker AuthChecker
}

// AddArticle can use any function compatible with AuthChecker interface (for example AuthModuleAuthChecker).
func (c AddArticle) AddArticle(username string, title string) error {
	if !c.AuthChecker.CheckAuth(username) {
		return fmt.Errorf("user %s is not allowed to add article", username)
	}

	// todo - some article adding logic
	fmt.Printf("added %s article by %s\n", title, username)

	return nil
}
