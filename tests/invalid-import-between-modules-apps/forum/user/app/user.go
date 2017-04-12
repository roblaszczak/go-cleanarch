package app

import "github.com/roblaszczak/go-cleanarch/tests/invalid-import-between-modules-apps/forum/post/app"

type User struct {
	Posts []app.Post
}