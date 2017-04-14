package main

import (
	"github.com/roblaszczak/go-cleanarch/examples/valid-cross-module-deps/cms/content/infrastructure"
	"github.com/roblaszczak/go-cleanarch/examples/valid-cross-module-deps/cms/content/usecases"
)

func main() {
	usecases.AddArticle{infrastructure.AuthModuleAuthChecker{}}.AddArticle("admin", "test article")
}
