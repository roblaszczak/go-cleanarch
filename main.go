package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/roblaszczak/go-cleanarch/cleanarch"
)

func main() {
	ignoredPackages := sliceFlag{}

	ignoreTests := flag.Bool("ignore-tests", false, "if flag is passed *_test.go files will be not checked")
	debug := flag.Bool("debug", false, "debug mode")
	flag.Var(
		&ignoredPackages,
		"ignore-package",
		"provided packages can be imported to any layer, "+
			"for example you can use`-ignore-package github.com/roblaszczak/go-cleanarch/infrastructure` to import "+
			"this package to the domain",
	)

	flag.Parse()
	var path string

	if *debug {
		cleanarch.Log.SetOutput(os.Stderr)
	}

	if len(flag.Args()) > 1 {
		path = flag.Args()[1]
	} else {
		var err error
		path, err = filepath.Abs(flag.Arg(0))
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf("[cleanarch] checking %s\n", path)

	validator := cleanarch.NewValidator()
	isValid, errors, err := validator.Validate(path, *ignoreTests, ignoredPackages)
	if err != nil {
		panic(err)
	}

	if !isValid {
		for _, err := range errors {
			fmt.Println(err.Error())
		}

		fmt.Println("Uncle Bob is not happy.")
		os.Exit(1)
	}

	os.Exit(0)
}
