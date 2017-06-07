package app

import "github.com/roblaszczak/go-cleanarch/examples/invalid-cross-module-deps/forum/post/app"

// User imports app.Post from another's module app. Only interfaces layer can be used in another modules.
type User struct {
	Posts []app.Post
}
