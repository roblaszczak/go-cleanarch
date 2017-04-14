package app

import "github.com/roblaszczak/go-cleanarch/examples/invalid-cross-module-deps/forum/post/app"

type User struct {
	Posts []app.Post
}
