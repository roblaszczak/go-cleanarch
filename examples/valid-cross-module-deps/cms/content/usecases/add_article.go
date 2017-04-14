package usecases

import (
	"fmt"
)

type AuthChecker interface {
	CheckAuth(username string) bool
}

type AddArticle struct {
	AuthChecker AuthChecker
}

func (c AddArticle) AddArticle(username string, title string) error {
	if !c.AuthChecker.CheckAuth(username) {
		return fmt.Errorf("user %s is not allowed to add article", username)
	}

	// todo - some article adding logic
	fmt.Printf("added %s article by %s\n", title, username)

	return nil
}
